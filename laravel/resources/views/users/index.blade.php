@extends('layouts.app')

@section('title', 'Data User')
@section('page-title', 'Data User')
@section('page-description', 'Kelola semua data user dan keluarga')

@section('page-actions')
    <a href="{{ route('users.create') }}" class="btn btn-light">
        <i class="fas fa-plus me-1"></i>Tambah User
    </a>
@endsection

@section('content')
<div class="card">
    <div class="card-header bg-white">
        <div class="row align-items-center">
            <div class="col-md-6">
                <h5 class="card-title mb-0">
                    <i class="fas fa-users me-2 text-primary"></i>Daftar User
                </h5>
            </div>
            <div class="col-md-6">
                <div class="input-group">
                    <input type="text" class="form-control" placeholder="Cari user..." id="searchUser">
                    <button class="btn btn-outline-primary" type="button">
                        <i class="fas fa-search"></i>
                    </button>
                </div>
            </div>
        </div>
    </div>
    <div class="card-body p-0">
        @if(isset($customers['data']) && count($customers['data']) > 0)
            <div class="table-responsive">
                <table class="table table-hover mb-0">
                    <thead>
                        <tr>
                            <th>No</th>
                            <th>Nama</th>
                            <th>Tanggal Lahir</th>
                            <th>Telepon</th>
                            <th>Email</th>
                            <th>Kewarganegaraan</th>
                            <th>Jumlah Keluarga</th>
                            <th>Aksi</th>
                        </tr>
                    </thead>
                    <tbody>
                        @foreach($customers['data'] as $index => $customer)
                            <tr>
                                <td>{{ $index + 1 }}</td>
                                <td>
                                    <div class="d-flex align-items-center">
                                        <div class="avatar bg-primary text-white rounded-circle me-3 d-flex align-items-center justify-content-center" style="width: 40px; height: 40px;">
                                            {{ strtoupper(substr($customer['cst_name'], 0, 1)) }}
                                        </div>
                                        <div>                                        <div class="fw-semibold">{{ $customer['cst_name'] }}</div>
                                        <small class="text-muted">ID: {{ $customer['cst_id'] }}</small>
                                        </div>
                                    </div>
                                </td>
                                <td>
                                    <span class="badge bg-light text-dark">
                                        {{ \Carbon\Carbon::parse($customer['cst_dob'])->format('d/m/Y') }}
                                    </span>
                                    <br>
                                    <small class="text-muted">
                                        ({{ \Carbon\Carbon::parse($customer['cst_dob'])->age }} tahun)
                                    </small>
                                </td>
                                <td>
                                    <i class="fas fa-phone me-1 text-success"></i>
                                    {{ $customer['cst_phoneNum'] }}
                                </td>
                                <td>
                                    <i class="fas fa-envelope me-1 text-info"></i>
                                    {{ $customer['cst_email'] }}
                                </td>
                                <td>
                                    @if(isset($nationalities['data']))
                                        @php
                                            $nationality = collect($nationalities['data'])->firstWhere('nationality_id', $customer['nationality_id']);
                                        @endphp
                                        @if($nationality)
                                            <span class="badge bg-primary">
                                                {{ $nationality['nationality_name'] }}
                                            </span>
                                        @else
                                            <span class="badge bg-secondary">Unknown</span>
                                        @endif
                                    @else
                                        <span class="badge bg-secondary">-</span>
                                    @endif
                                </td>
                                <td>
                                    <span class="badge bg-success">
                                        <i class="fas fa-users me-1"></i>
                                        {{ isset($customer['family_list']) ? count($customer['family_list']) : 0 }} orang
                                    </span>
                                </td>
                                <td>
                                    <div class="btn-group" role="group">
                                        <a href="{{ route('users.show', $customer['cst_id']) }}"
                                           class="btn btn-sm btn-outline-primary"
                                           title="Lihat Detail">
                                            <i class="fas fa-eye"></i>
                                        </a>
                                        <a href="{{ route('users.edit', $customer['cst_id']) }}"
                                           class="btn btn-sm btn-outline-warning"
                                           title="Edit">
                                            <i class="fas fa-edit"></i>
                                        </a>
                                        <form action="{{ route('users.destroy', $customer['cst_id']) }}"
                                              method="POST"
                                              class="d-inline"
                                              onsubmit="return confirm('Apakah Anda yakin ingin menghapus data ini?')">
                                            @csrf
                                            @method('DELETE')
                                            <button type="submit"
                                                    class="btn btn-sm btn-outline-danger"
                                                    title="Hapus">
                                                <i class="fas fa-trash"></i>
                                            </button>
                                        </form>
                                    </div>
                                </td>
                            </tr>
                        @endforeach
                    </tbody>
                </table>
            </div>
        @else
            <div class="text-center py-5">
                <div class="mb-3">
                    <i class="fas fa-users fa-3x text-muted"></i>
                </div>
                <h5 class="text-muted">Belum ada data user</h5>
                <p class="text-muted mb-3">Mulai dengan menambahkan user pertama Anda</p>
                <a href="{{ route('users.create') }}" class="btn btn-primary">
                    <i class="fas fa-plus me-1"></i>Tambah User Pertama
                </a>
            </div>
        @endif
    </div>
</div>

<!-- Statistics Cards -->
@if(isset($customers['data']) && count($customers['data']) > 0)
<div class="row mt-4">
    <div class="col-md-3">
        <div class="card text-center">
            <div class="card-body">
                <i class="fas fa-users fa-2x text-primary mb-2"></i>
                <h4 class="fw-bold">{{ count($customers['data']) }}</h4>
                <p class="text-muted mb-0">Total User</p>
            </div>
        </div>
    </div>
    <div class="col-md-3">
        <div class="card text-center">
            <div class="card-body">
                <i class="fas fa-globe fa-2x text-success mb-2"></i>
                <h4 class="fw-bold">{{ isset($nationalities['data']) ? count($nationalities['data']) : 0 }}</h4>
                <p class="text-muted mb-0">Negara</p>
            </div>
        </div>
    </div>
    <div class="col-md-3">
        <div class="card text-center">
            <div class="card-body">
                <i class="fas fa-heart fa-2x text-danger mb-2"></i>
                <h4 class="fw-bold">
                    @php
                        $totalFamily = 0;
                        if(isset($customers['data'])) {
                            foreach($customers['data'] as $customer) {
                                $totalFamily += isset($customer['family_count']) ? $customer['family_count'] : 0;
                            }
                        }
                    @endphp
                    {{ $totalFamily }}
                </h4>
                <p class="text-muted mb-0">Anggota Keluarga</p>
            </div>
        </div>
    </div>
    <div class="col-md-3">
        <div class="card text-center">
            <div class="card-body">
                <i class="fas fa-calendar fa-2x text-warning mb-2"></i>
                <h4 class="fw-bold">{{ date('Y') }}</h4>
                <p class="text-muted mb-0">Tahun Aktif</p>
            </div>
        </div>
    </div>
</div>
@endif
@endsection

@push('scripts')
<script>
$(document).ready(function() {
    // Search functionality
    $('#searchUser').on('keyup', function() {
        var value = $(this).val().toLowerCase();
        $('tbody tr').filter(function() {
            $(this).toggle($(this).text().toLowerCase().indexOf(value) > -1)
        });
    });

    // Auto refresh every 30 seconds
    setInterval(function() {
        if (!$('#searchUser').is(':focus')) {
            location.reload();
        }
    }, 30000);
});
</script>
@endpush
