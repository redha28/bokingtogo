<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\Log;

class UserController extends Controller
{
    private $apiBaseUrl = 'http://localhost:8080/api';

    public function index()
    {
        try {
            $customers = Http::get($this->apiBaseUrl . '/customers')->json();
            $nationalities = Http::get($this->apiBaseUrl . '/nationalities')->json();

            return view('users.index', compact('customers', 'nationalities'));
        } catch (\Exception $e) {
            Log::error('Error fetching data: ' . $e->getMessage());
            return view('users.index', ['customers' => [], 'nationalities' => []]);
        }
    }

    public function create()
    {
        try {
            $nationalities = Http::get($this->apiBaseUrl . '/nationalities')->json();
            return view('users.create', compact('nationalities'));
        } catch (\Exception $e) {
            Log::error('Error fetching nationalities: ' . $e->getMessage());
            return view('users.create', ['nationalities' => []]);
        }
    }    public function store(Request $request)
    {
        $request->validate([
            'cst_name' => 'required|string|max:255',
            'cst_dob' => 'required|date',
            'cst_phoneNum' => 'required|string|max:20',
            'cst_email' => 'required|email|max:255',
            'nationality_id' => 'required|integer',
            'families' => 'array',
            'families.*.fl_name' => 'string|max:255',
            'families.*.fl_dob' => 'date',
            'families.*.fl_relation' => 'string|max:100',
        ]);

        try {
            // Prepare family list data
            $familyList = [];
            if ($request->has('families') && is_array($request->families)) {
                foreach ($request->families as $family) {
                    if (!empty($family['fl_name']) && !empty($family['fl_dob']) && !empty($family['fl_relation'])) {
                        $familyList[] = [
                            'fl_dob' => $family['fl_dob'],
                            'fl_name' => $family['fl_name'],
                            'fl_relation' => $family['fl_relation']
                        ];
                    }
                }
            }

            // Create customer with family data
            $customerData = [
                'cst_dob' => $request->cst_dob,
                'cst_email' => $request->cst_email,
                'cst_name' => $request->cst_name,
                'cst_phoneNum' => $request->cst_phoneNum,
                'family_list' => $familyList,
                'nationality_id' => (int)$request->nationality_id,
            ];

            Log::info('Sending customer data:', $customerData);

            $customerResponse = Http::post($this->apiBaseUrl . '/customers', $customerData);

            Log::info('API Response Status:', ['status' => $customerResponse->status()]);
            Log::info('API Response Body:', ['body' => $customerResponse->body()]);

            if ($customerResponse->successful()) {
                return redirect()->route('users.index')->with('success', 'Data berhasil disimpan');
            } else {
                $errorMessage = 'Gagal menyimpan data customer';
                if ($customerResponse->json() && isset($customerResponse->json()['message'])) {
                    $errorMessage .= ': ' . $customerResponse->json()['message'];
                }
                return back()->withErrors(['error' => $errorMessage])->withInput();
            }
        } catch (\Exception $e) {
            Log::error('Error storing customer: ' . $e->getMessage());
            return back()->withErrors(['error' => 'Terjadi kesalahan saat menyimpan data: ' . $e->getMessage()])->withInput();
        }
    }

    public function show($id)
    {
        try {
            $customer = Http::get($this->apiBaseUrl . "/customers/{$id}")->json();

            // Get nationality data if nationality_id exists
            $nationality = null;
            if (isset($customer['data']['nationality_id'])) {
                $nationalityResponse = Http::get($this->apiBaseUrl . "/nationalities/{$customer['data']['nationality_id']}");
                if ($nationalityResponse->successful()) {
                    $nationality = $nationalityResponse->json();
                }
            }

            return view('users.show', compact('customer', 'nationality'));
        } catch (\Exception $e) {
            Log::error('Error fetching customer: ' . $e->getMessage());
            return redirect()->route('users.index')->withErrors(['error' => 'Data tidak ditemukan']);
        }
    }

    public function edit($id)
    {
        try {
            $customer = Http::get($this->apiBaseUrl . "/customers/{$id}")->json();
            $families = Http::get($this->apiBaseUrl . "/customers/{$id}/families")->json();
            $nationalities = Http::get($this->apiBaseUrl . '/nationalities')->json();

            return view('users.edit', compact('customer', 'families', 'nationalities'));
        } catch (\Exception $e) {
            Log::error('Error fetching customer for edit: ' . $e->getMessage());
            return redirect()->route('users.index')->withErrors(['error' => 'Data tidak ditemukan']);
        }
    }    public function update(Request $request, $id)
    {
        Log::info('Update request received:', [
            'id' => $id,
            'method' => $request->method(),
            'all_data' => $request->all()
        ]);

        $request->validate([
            'cst_name' => 'required|string|max:255',
            'cst_dob' => 'required|date',
            'cst_phoneNum' => 'required|string|max:20',
            'cst_email' => 'required|email|max:255',
            'nationality_id' => 'required|integer',
            'families' => 'array',
            'families.*.fl_name' => 'string|max:255',
            'families.*.fl_dob' => 'date',
            'families.*.fl_relation' => 'string|max:100',
        ]);

        try {
            // Prepare family list data
            $familyList = [];
            if ($request->has('families') && is_array($request->families)) {
                foreach ($request->families as $family) {
                    if (!empty($family['fl_name']) && !empty($family['fl_dob']) && !empty($family['fl_relation'])) {
                        $familyList[] = [
                            'fl_dob' => $family['fl_dob'],
                            'fl_name' => $family['fl_name'],
                            'fl_relation' => $family['fl_relation']
                        ];
                    }
                }
            }

            // Update customer with family data
            $customerData = [
                'cst_dob' => $request->cst_dob,
                'cst_email' => $request->cst_email,
                'cst_name' => $request->cst_name,
                'cst_phoneNum' => $request->cst_phoneNum,
                'family_list' => $familyList,
                'nationality_id' => (int)$request->nationality_id,
            ];

            Log::info('Sending UPDATE request to API:', [
                'url' => $this->apiBaseUrl . "/customers/{$id}",
                'method' => 'PUT',
                'data' => $customerData
            ]);

            $customerResponse = Http::put($this->apiBaseUrl . "/customers/{$id}", $customerData);

            Log::info('UPDATE API Response:', [
                'status' => $customerResponse->status(),
                'headers' => $customerResponse->headers(),
                'body' => $customerResponse->body()
            ]);

            if ($customerResponse->successful()) {
                return redirect()->route('users.index')->with('success', 'Data berhasil diupdate');
            } else {
                $errorMessage = 'Gagal mengupdate data customer';
                $responseData = $customerResponse->json();
                if ($responseData && isset($responseData['message'])) {
                    $errorMessage .= ': ' . $responseData['message'];
                }
                Log::error('Update failed:', [
                    'status' => $customerResponse->status(),
                    'response' => $responseData,
                    'error' => $errorMessage
                ]);
                return back()->withErrors(['error' => $errorMessage])->withInput();
            }
        } catch (\Exception $e) {
            Log::error('Error updating customer:', [
                'exception' => $e->getMessage(),
                'trace' => $e->getTraceAsString()
            ]);
            return back()->withErrors(['error' => 'Terjadi kesalahan saat mengupdate data: ' . $e->getMessage()])->withInput();
        }
    }

    public function destroy($id)
    {
        Log::info('DESTROY method called:', [
            'id' => $id,
            'request_method' => request()->method(),
            'url' => request()->url(),
            'all_data' => request()->all()
        ]);

        try {
            $response = Http::delete($this->apiBaseUrl . "/customers/{$id}");

            Log::info('DELETE API Response:', [
                'status' => $response->status(),
                'body' => $response->body()
            ]);

            if ($response->successful()) {
                return redirect()->route('users.index')->with('success', 'Data berhasil dihapus');
            } else {
                return back()->withErrors(['error' => 'Gagal menghapus data']);
            }
        } catch (\Exception $e) {
            Log::error('Error deleting customer: ' . $e->getMessage());
            return back()->withErrors(['error' => 'Terjadi kesalahan saat menghapus data']);
        }
    }
}
