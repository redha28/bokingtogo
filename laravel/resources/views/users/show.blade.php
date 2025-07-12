@extends('layouts.app')

@section('title', 'Detail User')
@section('page-title', 'Detail User')
@section('page-description', 'Informasi lengkap user dan keluarga')

@section('page-actions')
    <div class="d-flex gap-2 justify-content-end">
        <a href="{{ route('users.index') }}" class="btn btn-light">
            <i class="fas fa-arrow-left me-1"></i>Kembali
        </a>
    </div>
@endsection

@section('content')
@if(isset($customer['data']))
<div class="row">
    <!-- User Information Card -->
    <div class="col-md-8">
        <div class="card mb-4">
            <div class="card-header bg-white">
                <h5 class="card-title mb-0">
                    <i class="fas fa-user me-2 text-primary"></i>Informasi User
                </h5>
            </div>
            <div class="card-body">
                <div class="row">
                    <div class="col-md-6 mb-3">
                        <label class="fw-semibold text-muted">Nama Lengkap</label>
                        <p class="fs-5 fw-bold text-dark">{{ $customer['data']['cst_name'] }}</p>
                    </div>
                    <div class="col-md-6 mb-3">
                        <label class="fw-semibold text-muted">Tanggal Lahir</label>
                        <p class="fs-6">
                            {{ \Carbon\Carbon::parse($customer['data']['cst_dob'])->format('d F Y') }}
                            <span class="badge bg-light text-dark ms-2">
                                {{ \Carbon\Carbon::parse($customer['data']['cst_dob'])->age }} tahun
                            </span>
                        </p>
                    </div>
                    <div class="col-md-6 mb-3">
                        <label class="fw-semibold text-muted">Nomor Telepon</label>
                        <p class="fs-6">
                            <i class="fas fa-phone me-2 text-success"></i>
                            <a href="tel:{{ $customer['data']['cst_phoneNum'] }}" class="text-decoration-none">
                                {{ $customer['data']['cst_phoneNum'] }}
                            </a>
                        </p>
                    </div>
                    <div class="col-md-6 mb-3">
                        <label class="fw-semibold text-muted">Email</label>
                        <p class="fs-6">
                            <i class="fas fa-envelope me-2 text-info"></i>
                            <a href="mailto:{{ $customer['data']['cst_email'] }}" class="text-decoration-none">
                                {{ $customer['data']['cst_email'] }}
                            </a>
                        </p>
                    </div>
                    <div class="col-md-6 mb-3">
                        <label class="fw-semibold text-muted">Kewarganegaraan</label>
                        <p class="fs-6">
                            @if(isset($nationality['data']))
                                <span class="badge bg-primary fs-6">
                                    <i class="fas fa-flag me-1"></i>
                                    {{ $nationality['data']['nationality_name'] }} ({{ $nationality['data']['nationality_code'] }})
                                </span>
                            @else
                                <span class="badge bg-secondary fs-6">
                                    ID: {{ $customer['data']['nationality_id'] }}
                                </span>
                            @endif
                        </p>
                    </div>
                    <div class="col-md-6 mb-3">
                        <label class="fw-semibold text-muted">User ID</label>
                        <p class="fs-6">
                            <code>#{{ $customer['data']['cst_id'] }}</code>
                        </p>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <!-- User Avatar Card -->
    <div class="col-md-4">
        <div class="card mb-4">
            <div class="card-body text-center">
                <div class="avatar bg-primary text-white rounded-circle mx-auto mb-3 d-flex align-items-center justify-content-center" style="width: 100px; height: 100px; font-size: 2.5rem;">
                    {{ strtoupper(substr($customer['data']['cst_name'], 0, 1)) }}
                </div>
                <h5 class="fw-bold">{{ $customer['data']['cst_name'] }}</h5>
                <p class="text-muted">{{ $customer['data']['cst_email'] }}</p>

                <div class="d-grid gap-2">
                    <a href="{{ route('users.edit', $customer['data']['cst_id']) }}" class="btn btn-primary">
                        <i class="fas fa-edit me-1"></i>Edit Profile
                    </a>
                    <form action="{{ route('users.destroy', $customer['data']['cst_id']) }}"
                          method="POST"
                          onsubmit="return confirm('Apakah Anda yakin ingin menghapus data ini?')">
                        @csrf
                        @method('DELETE')
                        <button type="submit" class="btn btn-outline-danger w-100">
                            <i class="fas fa-trash me-1"></i>Hapus User
                        </button>
                    </form>
                </div>
            </div>
        </div>
    </div>
</div>

<!-- Family Members Section -->
<div class="card">
    <div class="card-header bg-white">
        <div class="d-flex justify-content-between align-items-center">
            <h5 class="card-title mb-0">
                <i class="fas fa-users me-2 text-primary"></i>Anggota Keluarga
            </h5>
            <span class="badge bg-primary">
                {{ isset($customer['data']['family_list']) ? count($customer['data']['family_list']) : 0 }} orang
            </span>
        </div>
    </div>
    <div class="card-body">
        @if(isset($customer['data']['family_list']) && count($customer['data']['family_list']) > 0)
            <div class="row">
                @foreach($customer['data']['family_list'] as $index => $family)
                    <div class="col-md-6 mb-3">
                        <div class="card border">
                            <div class="card-body">
                                <div class="d-flex align-items-center mb-3">
                                    <div class="avatar bg-success text-white rounded-circle me-3 d-flex align-items-center justify-content-center" style="width: 50px; height: 50px;">
                                        {{ strtoupper(substr($family['fl_name'], 0, 1)) }}
                                    </div>
                                    <div>
                                        <h6 class="fw-bold mb-0">{{ $family['fl_name'] }}</h6>
                                        <small class="text-muted">Anggota #{{ $index + 1 }}</small>
                                    </div>
                                </div>

                                <div class="row">
                                    <div class="col-12 mb-2">
                                        <small class="text-muted">Tanggal Lahir</small>
                                        <p class="mb-0 fw-semibold">
                                            {{ \Carbon\Carbon::parse($family['fl_dob'])->format('d/m/Y') }}
                                            <span class="badge bg-light text-dark ms-1">
                                                {{ \Carbon\Carbon::parse($family['fl_dob'])->age }} tahun
                                            </span>
                                        </p>
                                    </div>
                                    <div class="col-12">
                                        <small class="text-muted">Hubungan Keluarga</small>
                                        <p class="mb-0">
                                            <span class="badge bg-info">{{ $family['fl_relation'] }}</span>
                                        </p>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                @endforeach
            </div>
        @else
            <div class="text-center py-4">
                <i class="fas fa-users fa-3x text-muted mb-3"></i>
                <h6 class="text-muted">Belum ada anggota keluarga</h6>
                <p class="text-muted mb-3">User ini belum menambahkan data keluarga</p>
                <a href="{{ route('users.edit', $customer['data']['cst_id']) }}" class="btn btn-outline-primary">
                    <i class="fas fa-plus me-1"></i>Tambah Anggota Keluarga
                </a>
            </div>
        @endif
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
    // Add some interactive features
    $('.avatar').hover(
        function() {
            $(this).addClass('shadow-lg');
        },
        function() {
            $(this).removeClass('shadow-lg');
        }
    );
});
</script>
@endpush
