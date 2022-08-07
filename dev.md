#### Checklist Module

- [ ] FrontWeb
    - [ ] GetAllAssetsDefault elasticsearch pagination di frontend
    - [ ] GetAllAssetsSearch elasticsearch pagination di frontend
        - params:
            - property
                - name
                - lokasi
                - tipe
                - harga
            - otomotif
                - name
                - merk
                - seri
                - harga
            - mesin
                - name
                - lokasi
                - merk
                - harga
            - lainnya
                - name
                - lokasi
                - merk
    
- [ ] Admin 
    - [ ] Assets
        - [x] GetAll
        - [x] Store 
            type => 
            - form_b1| form bangunan 1: rumah, ruko, apartemen, condotel, kos atau kontrakan, toko, vila , kios
            - form_k1| form kendaraan 1: Kapal
            - form_k2| form kendaraan 2: Pesawat, Alat Berat
            - form_k3| from kendaraan 3: Mobil, Sepeda Motor, Truk, Bus
        - [x] GetOne
        - [x] Update
        - [x] Delete
    - [ ] BuildingAssets [repoonly]
        - [x] GetAll
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
    - [ ] VehicleAssets [repoonly]
        - [x] GetAll
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
    - [ ] Jadwal Lelang
        - [ ] GetAll
        - [ ] Store
        - [ ] GetOne
        - [ ] Update
        - [ ] Delete
    - [ ] Categories
        - [x] GetAll
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
    - [ ] SubCategories
        - [x] GetAll
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
    - [x] AssetAccessPlaces ketika create asset -> create ini [repoonly]
        - [x] GetAll
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
          - repository not yet change to *sql.DB
    - [ ] Facilities ketika create asset -> create ini [repoonly]
        - [x] GetAll
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
    - [ ] AssetFacilities ketika create asset -> create ini [repoonly]
        - [x] GetAll
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
    - [ ] Contacts ketika create asset -> create contact [repoonly]
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
    - [ ] Addresses ketika create asset -> create address [repoonly]
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
    - [ ] Images ketika create asset -> create images [repoonly]
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
    - [ ] AssetImages ketika create asset -> create images [repoonly]
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
    - [ ] Approvals ketika create asset -> create approvals [repoonly]
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
    - [ ] AssetApprovals ketika create asset -> create approvals [repoonly]
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
    - [ ] Kpknl
        - [x] GetAll
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
        - repository not yet change to *sql.DB
    - [ ] Faqs
        - [x] GetAll
        - [x] Store
        - [x] GetOne
        - [x] Update
        - [x] Delete
          - repository not yet change to *sql.DB
    - [ ] brands
        - [ ] GetAll
        - [ ] Store
        - [ ] GetOne
        - [ ] Update
        - [ ] Delete
    - [ ] sub_brands
        - [ ] GetAll
        - [ ] Store
        - [ ] GetOne
        - [ ] Update
        - [ ] Delete
    - [ ] type_of_certificate
        - [ ] GetAll
        - [ ] Store
        - [ ] GetOne
        - [ ] Update
        - [ ] Delete
    - [ ] machine_capacity
        - [ ] GetAll
        - [ ] Store
        - [ ] GetOne
        - [ ] Update
        - [ ] Delete
    - [ ] vehicle_colors
        - [ ] GetAll
        - [ ] Store
        - [ ] GetOne
        - [ ] Update
        - [ ] Delete
    - [ ] transmissions
        - [ ] GetAll
        - [ ] Store
        - [ ] GetOne
        - [ ] Update
        - [ ] Delete
    - [ ] loginUserByPN ldap mas tomi
    - [ ] create asset multiple form
    - [ ] handle upload witbh minio
    - [ ] scheduler buat update asset status to unpublish dan hapus asset di elastic

    
- Kpknl list
- category list
- sub category list where category_id
- postalcode like postalcode
- certificate_type
- facilities
- access_place
- checker list
- signer list

# Elastic trial 14Day
https://infolelang.es.us-central1.gcp.cloud.es.io
username,
elastic,
password 
CIEkxlGYVzr83ql1Lvwp4C1N

docker exec df2152ab9116 /usr/bin/mysqldump -u root --password=P@ssw0rd infolelang2 > infolelang2.sql 

-- ==========================================================
## PR Backend
-- auth onegate => done ,( incognito not done)
-- getMCS => done 
    -- - checker list
    -- - signer list

  - get data empty, belum semua di controller
  - pagination done
  - api inquiry cif ke esb = done

   - create vehicle done
        - certificate_type to certificate_type_id done
        - series_id done
        - brand_id done
        - transmission_id done
        - machineCapacity_id to machine_capacity_id
        - color_id done
        - get one vehicle join details above done

   - create to elastic done
   - create endpoint api delete image done
   - create api auction_schedule/ jadwal lelang, done query
   - api crud maintain banners done


 ### ORM
 - Tipe Update
 ```
 -- only include field to update
 func (asset AssetRepository) UpdateApproval(request *models.AssetsUpdateApproval, include []string) (responses bool, err error) {
	return true, asset.db.DB.Select(include).Updates(&request).Error

}

exclude := []string{
 	"last_maker_id",
 		"last_maker_desc",
 		"last_maker_date",
 		"published",
 		"publish_date",
 		"expired_date",
 }
-- Exclude field
func (asset AssetRepository) UpdateApproval(request *models.AssetsUpdateApproval,exclude []string) (responses bool, err error) {
     	return true, asset.db.DB.Omit(strings.Join(exclude, ","))Updates(&request).Error
 }

func (asset AssetRepository) UpdateApproval(request *models.AssetsUpdateApproval) (responses bool, err error) {
     	return true, asset.db.DB.Save(&request).Error
 }
 ``` 

   - create get all asset to elastic done endpoint , riset query
  # UNDONE  ====================================== 
   - create get all with filter to elastic done endpoint , riset query
      - filters
        1. properti
            - name
            - addresses.city
   
   - activity logger
   - runtime logger
   - handle get table error, 

   - api get menu by role , user matrix
        # get Menu done, by role not done
   - adding golang scheduler manager done
   - build to docker
   - prepare icon facilities, access places
   - handle getassetelastic 
     - access_place, facilities, bisa null di elastic
   - getOne ada data yg ga ke save, vehicle asset
   - encryptor response request AES 
   - nitip mas, endpoint updateApproval, pas update signer, loading nya lama, respon nya internal error, tp data nya masuk di elastic
   - yg getOne, kyknya data status akses nya, masih string 1 smua
   - get document id and update to table asset not handled

# next development
   - module user integration register with google or email,
   - create whitelist asset
   - mobile infolelang
   - notification created asset to approvals
   - integration geotaging google maps
   - get all status, name and desc
 ==========================================================

# New Stack
- Schema retry coding
  - https://stackoverflow.com/questions/67069723/keep-retrying-a-function-in-golang
  - https://upgear.io/blog/simple-golang-retry-function/
    - disetiap ada integrasi api keluar berikan retry


# Tunnel Server DEV
  - hostname: infolelang_admin.nopezi.com/infolelang
    service: http://192.168.1.7/infolelang
    
  - hostname: services.nopezi.com/infolelang_admin
    service: http://192.168.1.7/infolelang_admin
  
  - hostname: services.nopezi.com/pbb-api
    service: http://192.168.1.7/pbb-api
  
  - hostname: services.nopezi.com/pab
    service: http://192.168.1.7/pab
  
  - hostname: https://services.nopezi.com/elastic
    service: http://192.168.1.7/elastic
  
  - hostname: https://services.nopezi.com/kibana
    service: http://192.168.1.7/kibana
  
  - hostname: services.nopezi.com/minioapi
    service: http://192.168.1.7/minioapi

  - hostname: services.nopezi.com/minio
    service: http://192.168.1.7/minio

  - hostname: ssh.nopezi.com
    service: http://192.168.1.7:22

  - hostname: mysql.nopezi.com
    service: tcp://192.168.1.7:3306

chroot to docker mysql
docker exec -it mysql bash -l

minio 
minio1_1  | API: http://172.23.0.2:9000  http://127.0.0.1:9000 
minio1_1  | 
minio1_1  | Console: http://172.23.0.2:9001 http://127.0.0.1:9001 

```
##
# You should look at the following URL's in order to grasp a solid understanding
# of Nginx configuration files in order to fully unleash the power of Nginx.
# https://www.nginx.com/resources/wiki/start/
# https://www.nginx.com/resources/wiki/start/topics/tutorials/config_pitfalls/
# https://wiki.debian.org/Nginx/DirectoryStructure
#
# In most cases, administrators will remove this file from sites-enabled/ and
# leave it as reference inside of sites-available where it will continue to be
# updated by the nginx packaging team.
#
# This file will automatically load configuration files provided by other
# applications, such as Drupal or Wordpress. These applications will be made
# available underneath a path with that package name, such as /drupal8.
#
# Please see /usr/share/doc/nginx-doc/examples/ for more detailed examples.
##

# Default server configuration
#
server {
	listen 80 default_server;
	listen [::]:80 default_server;

	# SSL configuration
	#
	listen 443 ssl default_server;
	listen [::]:443 ssl default_server;
	#
	# Note: You should disable gzip for SSL traffic.
	# See: https://bugs.debian.org/773332
	#
	# Read up on ssl_ciphers to ensure a secure configuration.
	# See: https://bugs.debian.org/765782
	#
	# Self signed certs generated by the ssl-cert package
	# Don't use them in a production server!
	#
	# include snippets/snakeoil.conf;

	include snippets/self-signed.conf;
    	include snippets/ssl-params.conf;

	root /var/www/html;

	# Add index.php to the list if you are using PHP
	index index.html index.htm index.nginx-debian.html;

	server_name ccp13.dev.bri.co.id;

	location / {
		# First attempt to serve request as file, then
		# as directory, then fall back to displaying a 404.
		try_files $uri $uri/ =404;
	}

	location /elastic/ {
                proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
                proxy_set_header X-Forwarded-Proto $scheme;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_set_header Host $http_host;
                proxy_pass https://172.18.241.57:9200/;
        }

	location /kibana/ {
		#rewrite ^/minio(/.*)$ $1 break;
                proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
                proxy_set_header X-Forwarded-Proto $scheme;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_set_header Host $http_host;
                proxy_pass https://172.18.241.57:5601/;
        }

	location /api-eform-v3/ {
                proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
                proxy_set_header X-Forwarded-Proto $scheme;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_set_header Host $http_host;
                proxy_pass http://172.18.241.57:5999/;
        }
        location /portainer/ {
                proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
                proxy_set_header X-Forwarded-Proto $scheme;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_set_header Host $http_host;
                proxy_pass http://172.18.241.57:7000/;
        }

	location /minio/ {
#		rewrite ^minio/(.*)$ $1;
		rewrite ^/minio(/.*)$ $1 break;
		#rewrite ^/minio(.*) /$1 break;
                proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
                proxy_set_header X-Forwarded-Proto $scheme;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_set_header Host $http_host;
                proxy_pass http://localhost:9001;
        }

        location /minioapi/ {
                proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
                proxy_set_header X-Forwarded-Proto $scheme;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_set_header Host $http_host;
                proxy_pass http://localhost:9000/;
        }



        location /infolelang/ {
                proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
                proxy_set_header X-Forwarded-Proto $scheme;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_set_header Host $http_host;
                proxy_pass http://172.18.241.57:5000/;
        }

        #location /pab_admin/ {
        #        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        #        proxy_set_header X-Forwarded-Proto $scheme;
        #        proxy_set_header X-Real-IP $remote_addr;
        #        proxy_set_header Host $http_host;
        #        proxy_pass http://172.18.241.57:3000/;
        #}

        location /infolelang_admin {
                alias /var/www/html/infolelang_admin/dist;

                try_files $uri $uri/ @infolelang_admin;
        }

        location @infolelang_admin {
                rewrite /infolelang_admin/(.*)$ /infolelang_admin/index.html last;
        }

        location /pab {
                alias /var/www/html/pab_admin_web/dist;

                try_files $uri $uri/ @pbb;
        }

        location @pbb {
                rewrite /pbb_admin/(.*)$ /pbb_admin/index.html last;
        }

	location /pbb-api/ {
               proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
               proxy_set_header X-Forwarded-Proto $scheme;
               proxy_set_header X-Real-IP $remote_addr;
               proxy_set_header Host $http_host;
               proxy_pass http://172.18.241.57:2243/;
	}

        location /infolelang_web {
                alias /var/www/html/infolelang_web/dist;

                try_files $uri $uri/ @infolelang_web;
        }

        location @infolelang_web {
                rewrite /infolelang_web/(.*)$ /infolelang_web/index.html last;
        }

	# pass PHP scripts to FastCGI server
	#
	#location ~ \.php$ {
	#	include snippets/fastcgi-php.conf;
	#
	#	# With php-fpm (or other unix sockets):
	#	fastcgi_pass unix:/var/run/php/php7.4-fpm.sock;
	#	# With php-cgi (or other tcp sockets):
	#	fastcgi_pass 127.0.0.1:9000;
	#}

	# deny access to .htaccess files, if Apache's document root
	# concurs with nginx's one
	#
	#location ~ /\.ht {
	#	deny all;
	#}
}


# Virtual Host configuration for example.com
#
# You can move that to a different file under sites-available/ and symlink that
# to sites-enabled/ to enable it.
#
#server {
#	listen 80;
#	listen [::]:80;
#
#	server_name example.com;
#
#	root /var/www/example.com;
#	index index.html;
#
#	location / {
#		try_files $uri $uri/ =404;
#	}
#}

```


# MINIO
nano  /etc/systemd/system/minio.service

[Unit]
Description=MinIO
Documentation=https://docs.min.io
Wants=network-online.target
After=network-online.target
AssertFileIsExecutable=/usr/local/bin/minio

[Service]
WorkingDirectory=/usr/local/

User=minio-user
Group=minio-user
ProtectProc=invisible

EnvironmentFile=/etc/default/minio
ExecStartPre=/bin/bash -c "if [ -z \"${MINIO_VOLUMES}\" ]; then echo \"Variable MINIO_VOLUMES not set in /etc/default/minio\"; exit 1; fi"
ExecStart=/usr/local/bin/minio server --console-address :33639 $MINIO_OPTS $MINIO_VOLUMES

# Let systemd restart this service always
Restart=always

# Specifies the maximum file descriptor number that can be opened by this process
LimitNOFILE=1048576

# Specifies the maximum number of threads this process can create
TasksMax=infinity

# Disable timeout logic and wait until process is stopped
TimeoutStopSec=infinity
SendSIGKILL=no

[Install]
WantedBy=multi-user.target

# Built for ${project.name}-${project.version} (${project.name})


root@ubuntu-etilangv2:~# cat /etc/default/minio
MINIO_ACCESS_KEY="minio"
MINIO_VOLUMES="/usr/local/share/minio/"
MINIO_OPTS="-C /etc/minio --address 128.199.194.80:9000"  "/minio"
MINIO_SECRET_KEY="P@ssw0rd"


## BUG API
- aku coba, tolak nya dah bisa, dari tolak checker, statusnya brubah ditolak checker, cman api nya kudu di reload dlu, baru lancar
- dari approver checker, statusnya gk brubah ke pending signer

- expired date diisi saat apa ?

==============================
- integrasi api mas tomi, mcs, login
- nambah uker ke table asset dari data login
- integrasi esb buat ambildata cif
- user matrix database,
- login by onegate, 
- get user by onegate
- get menu by jabatan

update publish signer, status ilang di getall maintain done
dmin publish, unpublish, delete, update edit gaperlu cs, done
maker publish, unpublish perlu cs, delete  perlu cs done

### prioritas 2
updateMaintain
- status hilang done
- address ga ke update done
- buildingAsset ilang
- vehicleAsset ilang
- facilities
- aksesPlace
- contact
- document_id done
- images bisa
-  approvals ga ke update

## prioritas 1 done
- getAllApproval done | nanti abis ada user metrik dicek lg
 - published di by pn checker dan by pn signer = published dikosongin
 - pending checker by pn masing2,  pending signer by pn masing2 
 - deleted (engga jadi)

 ## get file minio lewat backend

 ## cek jwt kalo muti app bisa di check valid ga
 ## cryoto response dan request

 ## di get elastic datanya batasin
 hilangkan:
   - data nasabah
   - data approval