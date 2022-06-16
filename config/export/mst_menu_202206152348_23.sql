INSERT INTO infolelang2.mst_menu (Title,Url,Deskripsi,Icon,Atribut,Badge,IDParent,Target,Urutan,RoleAccess,KanpusOnly,Jenis,Posisi,Status) VALUES 
('Beranda','/','Beranda','mdi-view-dashboard',NULL,0,0,NULL,1,1,0,0,2,1)
,('Usulan PO','','Pengajuan PO','mdi-clipboard-list-outline',NULL,0,0,NULL,2,1,0,0,2,1)
,('Realisasi PO','','Realisasi PO','mdi-human-edit',NULL,0,0,NULL,3,1,0,0,2,1)
,('Report & Monitoring PO','/compliance/monitoring','Monitoring PO','mdi-cog-outline',NULL,0,0,NULL,4,1,0,0,2,1)
,('Usulan Compliance','','Monitoring PO','mdi-clipboard-list-outline',NULL,0,0,NULL,5,1,0,0,2,1)
,('Realisasi Compliance','','Realisasi Compliance','mdi-human-edit',NULL,0,0,NULL,6,1,0,0,2,1)
,('Report & Monitoring','/compliance/monitoring','Monitoring Compliance','mdi-cog-outline',NULL,0,0,NULL,7,1,0,0,2,1)
,('Buat Usulan','/produk-owner/pengajuan/create','Create',' ',NULL,0,2,NULL,1,1,0,0,2,1)
,('Draft Approval','/produk-owner/pengajuan/drafts','Draft Approval','fas fa-pencil-alt fa-fw',NULL,0,2,NULL,2,1,0,0,2,1)
,('Approval Usulan','/produk-owner/pengajuan/approvals','Approval Pengajuan',' ',NULL,0,2,NULL,3,1,0,0,2,1)
;
INSERT INTO infolelang2.mst_menu (Title,Url,Deskripsi,Icon,Atribut,Badge,IDParent,Target,Urutan,RoleAccess,KanpusOnly,Jenis,Posisi,Status) VALUES 
('Pending Realisasi','/produk-owner/realisasi/pendings','Pending Realisasi','fas fa-pencil-alt fa-fw',NULL,0,3,NULL,1,1,0,0,2,1)
,('Draft Realisasi','/produk-owner/realisasi/drafts','Draft Realisasi','fa fa-folder-open  fa-fw',NULL,0,3,NULL,2,1,0,0,2,1)
,('Approval Realisasi','/produk-owner/realisasi/approvals','Approval Realisasi','fa fa-folder-open  fa-fw',NULL,0,3,NULL,3,1,0,0,2,1)
,('Usulan Masuk','/compliance/pengajuan/inbox','Pengajuan Masuk','fas fa-pencil-alt fa-fw',NULL,0,5,NULL,1,1,0,0,2,1)
,('Verifikasi Usulan','/compliance/pengajuan/verifications','Verifikasi Pengajuan','fas fa-pencil-alt fa-fw',NULL,0,5,NULL,2,1,0,0,2,1)
,('Approval Verifikasi Usulan','/compliance/pengajuan/approvals','Approval Pengajuan','',NULL,0,5,NULL,4,1,0,0,2,1)
,('Realisasi Masuk','/compliance/realisasi/inbox','Realisasi Masuk',' ',NULL,0,6,NULL,1,1,0,0,2,1)
,('Verifikasi Realisasi','/compliance/realisasi/verifications','Verifikasi Realisasi',' ',NULL,0,6,NULL,2,1,0,0,2,1)
,('Approval Realisasi','/compliance/realisasi/approvals','Approval Realisasi',' ',NULL,0,6,NULL,3,1,0,0,2,1)
,('Persetujuan Regulator','/compliance/realisasi/persetujuan-regulator','Persetujuan Regulator',' ',NULL,0,5,NULL,3,1,0,0,2,1)
;