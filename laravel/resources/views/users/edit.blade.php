@extends('layouts.app')

@section('title', 'Edit User')
@section('page-title', 'Edit User')
@section('page-description', 'Ubah data user dan keluarganya')

@section('page-actions')
    <a href="{{ route('users.show', $customer['data']['cst_id']) }}" class="btn btn-light">
        <i class="fas fa-arrow-left me-1"></i>Kembali
    </a>
@endsection

@section('content')
@if(isset($customer['data']))
<div class="card">
    <div class="card-header bg-white">
        <h5 class="card-title mb-0">
            <i class="fas fa-user-edit me-2 text-primary"></i>Edit Data User: {{ $customer['data']['cst_name'] }}
        </h5>
    </div>
    <div class="card-body">
        <form action="{{ route('users.update', $customer['data']['cst_id']) }}" method="POST" id="editUserForm">
            @csrf
            @method('PUT')

            <!-- Data User Section -->
            <div class="row mb-4">
                <div class="col-12">
                    <h6 class="text-primary fw-bold mb-3">
                        <i class="fas fa-user me-2"></i>DATA USER
                    </h6>
                </div>

                <div class="col-md-6 mb-3">
                    <label for="cst_name" class="form-label fw-semibold">Nama <span class="text-danger">*</span></label>
                    <input type="text"
                           class="form-control @error('cst_name') is-invalid @enderror"
                           id="cst_name"
                           name="cst_name"
                           value="{{ old('cst_name', $customer['data']['cst_name']) }}"
                           placeholder="Masukkan nama anda"
                           required>
                    @error('cst_name')
                        <div class="invalid-feedback">{{ $message }}</div>
                    @enderror
                </div>

                <div class="col-md-6 mb-3">
                    <label for="cst_dob" class="form-label fw-semibold">Tanggal Lahir <span class="text-danger">*</span></label>
                    <input type="date"
                           class="form-control @error('cst_dob') is-invalid @enderror"
                           id="cst_dob"
                           name="cst_dob"
                           value="{{ old('cst_dob', $customer['data']['cst_dob']) }}"
                           required>
                    @error('cst_dob')
                        <div class="invalid-feedback">{{ $message }}</div>
                    @enderror
                </div>

                <div class="col-md-6 mb-3">
                    <label for="cst_phoneNum" class="form-label fw-semibold">Nomor Telepon <span class="text-danger">*</span></label>
                    <input type="tel"
                           class="form-control @error('cst_phoneNum') is-invalid @enderror"
                           id="cst_phoneNum"
                           name="cst_phoneNum"
                           value="{{ old('cst_phoneNum', $customer['data']['cst_phoneNum']) }}"
                           placeholder="Contoh: 08123456789"
                           required>
                    @error('cst_phoneNum')
                        <div class="invalid-feedback">{{ $message }}</div>
                    @enderror
                </div>

                <div class="col-md-6 mb-3">
                    <label for="cst_email" class="form-label fw-semibold">Email <span class="text-danger">*</span></label>
                    <input type="email"
                           class="form-control @error('cst_email') is-invalid @enderror"
                           id="cst_email"
                           name="cst_email"
                           value="{{ old('cst_email', $customer['data']['cst_email']) }}"
                           placeholder="contoh@email.com"
                           required>
                    @error('cst_email')
                        <div class="invalid-feedback">{{ $message }}</div>
                    @enderror
                </div>

                <div class="col-md-6 mb-3">
                    <label for="nationality_id" class="form-label fw-semibold">Kewarganegaraan <span class="text-danger">*</span></label>
                    <select class="form-select @error('nationality_id') is-invalid @enderror"
                            id="nationality_id"
                            name="nationality_id"
                            required>
                        <option value="">Pilih kewarganegaraan</option>
                        @if(isset($nationalities['data']) && is_array($nationalities['data']))
                            @foreach($nationalities['data'] as $nationality)
                                <option value="{{ $nationality['nationality_id'] }}"
                                        {{ old('nationality_id', $customer['data']['nationality_id']) == $nationality['nationality_id'] ? 'selected' : '' }}>
                                    {{ $nationality['nationality_name'] }} ({{ $nationality['nationality_code'] }})
                                </option>
                            @endforeach
                        @endif
                    </select>
                    @error('nationality_id')
                        <div class="invalid-feedback">{{ $message }}</div>
                    @enderror
                </div>
            </div>

            <hr>

            <!-- Data Keluarga Section -->
            <div class="row mb-4">
                <div class="col-12">
                    <div class="d-flex justify-content-between align-items-center mb-3">
                        <h6 class="text-primary fw-bold mb-0">
                            <i class="fas fa-users me-2"></i>KELUARGA
                        </h6>
                        <button type="button" class="btn btn-outline-primary btn-sm" id="addFamily">
                            <i class="fas fa-plus me-1"></i>Tambah Keluarga
                        </button>
                    </div>
                </div>
            </div>

            <div id="familyContainer">
                @if(isset($customer['data']['family_list']) && count($customer['data']['family_list']) > 0)
                    @foreach($customer['data']['family_list'] as $index => $family)
                        <div class="family-member p-3 mb-3">
                            <div class="d-flex justify-content-between align-items-center mb-3">
                                <h6 class="mb-0 text-secondary">
                                    <i class="fas fa-user me-2"></i>Anggota Keluarga #{{ $index + 1 }}
                                </h6>
                                <button type="button" class="btn btn-danger btn-sm remove-family">
                                    <i class="fas fa-trash"></i>
                                </button>
                            </div>

                            <div class="row">
                                <div class="col-md-6 mb-3">
                                    <label class="form-label fw-semibold">Nama <span class="text-danger">*</span></label>
                                    <input type="text"
                                           class="form-control"
                                           name="families[{{ $index }}][fl_name]"
                                           value="{{ $family['fl_name'] }}"
                                           placeholder="Masukkan Nama"
                                           required>
                                </div>

                                <div class="col-md-6 mb-3">
                                    <label class="form-label fw-semibold">Tanggal Lahir <span class="text-danger">*</span></label>
                                    <input type="date"
                                           class="form-control"
                                           name="families[{{ $index }}][fl_dob]"
                                           value="{{ $family['fl_dob'] }}"
                                           required>
                                </div>

                                <div class="col-md-6 mb-3">
                                    <label class="form-label fw-semibold">Hubungan Keluarga <span class="text-danger">*</span></label>
                                    <select class="form-select" name="families[{{ $index }}][fl_relation]" required>
                                        <option value="">Pilih hubungan</option>
                                        <option value="Suami" {{ $family['fl_relation'] == 'Suami' ? 'selected' : '' }}>Suami</option>
                                        <option value="Istri" {{ $family['fl_relation'] == 'Istri' ? 'selected' : '' }}>Istri</option>
                                        <option value="Anak" {{ $family['fl_relation'] == 'Anak' ? 'selected' : '' }}>Anak</option>
                                        <option value="Ayah" {{ $family['fl_relation'] == 'Ayah' ? 'selected' : '' }}>Ayah</option>
                                        <option value="Ibu" {{ $family['fl_relation'] == 'Ibu' ? 'selected' : '' }}>Ibu</option>
                                        <option value="Saudara" {{ $family['fl_relation'] == 'Saudara' ? 'selected' : '' }}>Saudara</option>
                                        <option value="Lainnya" {{ $family['fl_relation'] == 'Lainnya' ? 'selected' : '' }}>Lainnya</option>
                                    </select>
                                </div>
                            </div>
                        </div>
                    @endforeach
                @endif
            </div>

            <!-- Submit Button -->
            <div class="row mt-4">
                <div class="col-12">
                    <div class="d-flex gap-2">
                        <button type="submit" class="btn btn-primary">
                            <i class="fas fa-save me-1"></i>Update Data
                        </button>
                        <a href="{{ route('users.show', $customer['data']['cst_id']) }}" class="btn btn-outline-secondary">
                            <i class="fas fa-times me-1"></i>Batal
                        </a>
                    </div>
                </div>
            </div>
        </form>

        <!-- Separate Delete Form -->
        <div class="mt-3">
            <form action="{{ route('users.destroy', $customer['data']['cst_id']) }}"
                  method="POST"
                  class="d-inline"
                  onsubmit="return confirm('Apakah Anda yakin ingin menghapus data ini?')">
                @csrf
                @method('DELETE')
                <button type="submit" class="btn btn-danger">
                    <i class="fas fa-trash me-1"></i>Hapus User
                </button>
            </form>
        </div>
    </div>
</div>
@else
<div class="alert alert-danger">
    <i class="fas fa-exclamation-triangle me-2"></i>Data user tidak ditemukan.
</div>
@endif
@endsection

@push('scripts')
<script>
$(document).ready(function() {
    let familyIndex = {{ isset($customer['data']['family_list']) ? count($customer['data']['family_list']) : 0 }};

    // Add family member
    $('#addFamily').click(function() {
        addFamilyMember();
    });

    // Remove family member
    $(document).on('click', '.remove-family', function() {
        $(this).closest('.family-member').remove();
        updateFamilyNumbers();
    });

    function addFamilyMember() {
        const familyHtml = `
            <div class="family-member p-3 mb-3">
                <div class="d-flex justify-content-between align-items-center mb-3">
                    <h6 class="mb-0 text-secondary">
                        <i class="fas fa-user me-2"></i>Anggota Keluarga #${familyIndex + 1}
                    </h6>
                    <button type="button" class="btn btn-danger btn-sm remove-family">
                        <i class="fas fa-trash"></i>
                    </button>
                </div>

                <div class="row">
                    <div class="col-md-6 mb-3">
                        <label class="form-label fw-semibold">Nama <span class="text-danger">*</span></label>
                        <input type="text"
                               class="form-control"
                               name="families[${familyIndex}][fl_name]"
                               placeholder="Masukkan Nama"
                               required>
                    </div>

                    <div class="col-md-6 mb-3">
                        <label class="form-label fw-semibold">Tanggal Lahir <span class="text-danger">*</span></label>
                        <input type="date"
                               class="form-control"
                               name="families[${familyIndex}][fl_dob]"
                               required>
                    </div>

                    <div class="col-md-6 mb-3">
                        <label class="form-label fw-semibold">Hubungan Keluarga <span class="text-danger">*</span></label>
                        <select class="form-select" name="families[${familyIndex}][fl_relation]" required>
                            <option value="">Pilih hubungan</option>
                            <option value="Suami">Suami</option>
                            <option value="Istri">Istri</option>
                            <option value="Anak">Anak</option>
                            <option value="Ayah">Ayah</option>
                            <option value="Ibu">Ibu</option>
                            <option value="Saudara">Saudara</option>
                            <option value="Lainnya">Lainnya</option>
                        </select>
                    </div>
                </div>
            </div>
        `;

        $('#familyContainer').append(familyHtml);
        familyIndex++;
        updateFamilyNumbers();
    }

    function updateFamilyNumbers() {
        $('.family-member').each(function(index) {
            $(this).find('h6').html(`<i class="fas fa-user me-2"></i>Anggota Keluarga #${index + 1}`);
        });
    }

    // Form validation with more specific error logging
    $('#editUserForm').on('submit', function(e) {
        console.log('Form submission started');
        console.log('Form action:', $(this).attr('action'));
        console.log('Form method:', $(this).attr('method'));

        let isValid = true;

        // Check required fields
        $(this).find('[required]').each(function() {
            if (!$(this).val()) {
                isValid = false;
                $(this).addClass('is-invalid');
                console.log('Invalid field:', $(this).attr('name'));
            } else {
                $(this).removeClass('is-invalid');
            }
        });

        if (!isValid) {
            e.preventDefault();
            alert('Mohon lengkapi semua field yang wajib diisi (*)');
            return false;
        }

        console.log('Form validation passed, submitting...');
    });
});
</script>
@endpush
