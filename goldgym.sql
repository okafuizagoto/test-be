-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 15 Okt 2023 pada 19.09
-- Versi server: 10.4.21-MariaDB
-- Versi PHP: 8.0.11

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `goldgym`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `data_peserta`
--

CREATE TABLE `data_peserta` (
  `gold_id` int(11) NOT NULL,
  `gold_email` varchar(250) DEFAULT NULL,
  `gold_password` varchar(250) DEFAULT NULL,
  `gold_nama` varchar(250) DEFAULT NULL,
  `gold_nomorhp` varchar(250) DEFAULT NULL,
  `gold_nomorkartu` varchar(250) DEFAULT NULL,
  `gold_cvv` varchar(250) DEFAULT NULL,
  `gold_expireddate` datetime DEFAULT NULL,
  `gold_namapemegangkartu` varchar(350) DEFAULT NULL,
  `gold_validasiyn` char(1) NOT NULL DEFAULT 'N',
  `gold_token` varchar(250) DEFAULT NULL,
  `gold_otp` varchar(6) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `data_peserta`
--

INSERT INTO `data_peserta` (`gold_id`, `gold_email`, `gold_password`, `gold_nama`, `gold_nomorhp`, `gold_nomorkartu`, `gold_cvv`, `gold_expireddate`, `gold_namapemegangkartu`, `gold_validasiyn`, `gold_token`, `gold_otp`) VALUES
(0, 'testings', 'testing', 'testing', '0852312521', 'argon2$4$32768$4$32$13ALAH8O7r9Cywv4Z+ZyyQ==$yly2UQccckmdgx5wfkOSOslyp+S2E5xZmEf0CiilTUA=', 'arg', '2023-09-28 00:00:00', 'testing', 'N', NULL, NULL),
(1, 'test', 'testing', 'testing', '0852312521', '23582398', '244', '2023-09-28 00:00:00', 'testing', 'N', NULL, NULL),
(4, 'testingss', 'testing123', 'tester', '0852312521', 'argon2$4$32768$4$32$vjny9OU/VHsALna8j0v7LQ==$KnYnJCX+PCrSx7EG2mtlCqyFlUsK0tBAXPh5cLaNfr0=', 'argon2$4$32768$4$32$VL92PFobFcp086xHlQEGzA==$hWpG2jf0Pkx/iVpdxKRmWgPTNc3n/+iAXSilvafBraY=', '2023-09-28 00:00:00', 'argon2$4$32768$4$32$3LmUAVGaqMB403EpAe8ozw==$TwVGMZrGGJfsAbw2iI4TdFytY1m6sL7OPAas6aPmoL4=', 'N', NULL, NULL),
(5, 'okafuiz@gmail.com', 'testing123', 'tester', '0852312521', 'argon2$4$32768$4$32$vjny9OU/VHsALna8j0v7LQ==$KnYnJCX+PCrSx7EG2mtlCqyFlUsK0tBAXPh5cLaNfr0=', 'argon2$4$32768$4$32$VL92PFobFcp086xHlQEGzA==$hWpG2jf0Pkx/iVpdxKRmWgPTNc3n/+iAXSilvafBraY=', '2023-09-28 00:00:00', 'argon2$4$32768$4$32$3LmUAVGaqMB403EpAe8ozw==$TwVGMZrGGJfsAbw2iI4TdFytY1m6sL7OPAas6aPmoL4=', 'N', 'abcde12345', NULL);

-- --------------------------------------------------------

--
-- Struktur dari tabel `data_token`
--

CREATE TABLE `data_token` (
  `gold_token` varchar(250) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `data_token`
--

INSERT INTO `data_token` (`gold_token`) VALUES
('abcde12345');

-- --------------------------------------------------------

--
-- Struktur dari tabel `subscription`
--

CREATE TABLE `subscription` (
  `gold_id` int(11) NOT NULL,
  `gold_totalharga` float DEFAULT NULL,
  `gold_validasipayment` char(1) NOT NULL DEFAULT 'N',
  `gold_otp` varchar(6) DEFAULT NULL,
  `gold_lastupdate` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `subscription`
--

INSERT INTO `subscription` (`gold_id`, `gold_totalharga`, `gold_validasipayment`, `gold_otp`, `gold_lastupdate`) VALUES
(1, NULL, 'N', NULL, NULL),
(5, 150000, 'N', '406833', '2023-10-16 00:04:14');

-- --------------------------------------------------------

--
-- Struktur dari tabel `subscription_detail`
--

CREATE TABLE `subscription_detail` (
  `gold_id` int(11) NOT NULL,
  `gold_menuid` int(11) NOT NULL,
  `gold_namapaket` varchar(250) DEFAULT NULL,
  `gold_namalayanan` varchar(250) DEFAULT NULL,
  `gold_harga` float DEFAULT NULL,
  `gold_jadwal` varchar(250) DEFAULT NULL,
  `gold_listlatihan` varchar(250) DEFAULT NULL,
  `gold_jumlahpertemuan` int(11) DEFAULT NULL,
  `gold_durasi` int(11) DEFAULT NULL,
  `gold_startdate` datetime DEFAULT NULL,
  `gold_enddate` datetime DEFAULT NULL,
  `gold_statuslangganan` varchar(250) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `subscription_detail`
--

INSERT INTO `subscription_detail` (`gold_id`, `gold_menuid`, `gold_namapaket`, `gold_namalayanan`, `gold_harga`, `gold_jadwal`, `gold_listlatihan`, `gold_jumlahpertemuan`, `gold_durasi`, `gold_startdate`, `gold_enddate`, `gold_statuslangganan`) VALUES
(1, 1, 'testing', 'testing', 50000, 'senin-jumat', 'test', 13, 65, NULL, NULL, 'test'),
(1, 2, NULL, 'test', 80000, 'senin-jumat', 'testing', 12, 70, NULL, NULL, 'testing'),
(5, 1, 'test', 'test', 50000, 'senin,rabu,jumat', 'push up', 12, 60, NULL, NULL, 'Belum Berlangganan'),
(5, 2, 'testing', 'testing', 100000, 'selasa,kamis,sabtu', 'pull up', 12, 60, NULL, NULL, 'Belum Berlangganan');

-- --------------------------------------------------------

--
-- Struktur dari tabel `subscription_product`
--

CREATE TABLE `subscription_product` (
  `gold_menuid` int(11) NOT NULL,
  `gold_namapaket` varchar(250) DEFAULT NULL,
  `gold_namalayanan` varchar(250) DEFAULT NULL,
  `gold_harga` float DEFAULT NULL,
  `gold_jadwal` varchar(250) DEFAULT NULL,
  `gold_listlatihan` varchar(250) DEFAULT NULL,
  `gold_jumlahpertemuan` varchar(250) DEFAULT NULL,
  `gold_durasi` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `subscription_product`
--

INSERT INTO `subscription_product` (`gold_menuid`, `gold_namapaket`, `gold_namalayanan`, `gold_harga`, `gold_jadwal`, `gold_listlatihan`, `gold_jumlahpertemuan`, `gold_durasi`) VALUES
(1, 'test', 'test', 50000, 'senin,rabu,jumat', 'push up', '12', 60),
(2, 'testing', 'testing', 100000, 'selasa,kamis,sabtu', 'pull up', '12', 60);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `data_peserta`
--
ALTER TABLE `data_peserta`
  ADD PRIMARY KEY (`gold_id`);

--
-- Indeks untuk tabel `subscription`
--
ALTER TABLE `subscription`
  ADD PRIMARY KEY (`gold_id`);

--
-- Indeks untuk tabel `subscription_detail`
--
ALTER TABLE `subscription_detail`
  ADD PRIMARY KEY (`gold_id`,`gold_menuid`);

--
-- Indeks untuk tabel `subscription_product`
--
ALTER TABLE `subscription_product`
  ADD PRIMARY KEY (`gold_menuid`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
