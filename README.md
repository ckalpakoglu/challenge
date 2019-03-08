# FrontEnd Challenge-1

Örnek sunucuya username:password ikilisi atılarak JWT token alınmalı ve ardından
"sürükle bırak" olarak çalışacak ekteki .xml dosyasının yüklenmesini (ve görüntülenmesini)
beklenmektedir.

###### Not: Dosya yüklenir yüklenmez verinin "parse edilmiş" hali response içeriğinde dönecektir. Bu verinin bir data-grid yardımı ile gösterilmesi beklenmektedir.
###### Not: Data-Grid üzerinde filtreler çalıştırılabilir.
###### Not: Docker mimarisinde yapılması tercih sebebidir.

### Sunucu 3 adet endpoint içermektedir:
  * Unauthenticated Access:
    - "/login"
      - POST
    - "/"
      - GET
    - "health-check"
      - GET
  * Authenticated Access:
    - "/restricted"
      - GET
      - POST

### Request örnekleri:
HealthCheck:
      `curl -X GET localhost:8086/health-check`
Root:
      `curl -X GET localhost:8086/`
Login:
      `curl -X POST -d '{"username":"admin", "password":"admin123!"}' "localhost:8086/login" -H "Content-Type: application/json"`
HealthCheck:
      `curl -X GET localhost:8086/health-check`
Restricted POST (XML dosyası yükleme):
    `curl -XPOST -H "Authorization: Bearer <jwt_token>" -H "Content-Type: text/xml" --data @example-2C-scan-export-Burp.xml localhost:8086/restricted`
Restricted GET:
      `curl -XPOST -H "Authorization: Bearer <jwt_token>" localhost:8086/restricted`


### Compile & Run:
Uygulamayı cmd üzerinden çalıştırmak için:
     `$> ./challenge`

Eğer uygulama senin ortamında çalışmazsa compile/derleme işlemi yapılması geeklidir. Compile işlemi için GO ortamı kurulduktan sonra:

     $> go build && ./challenge

yapılması gereklidir.

Windows ortamında çalıştırmak için, cmd açıp ilgili dizine gittikten sonra:
    `$> challenge.exe`


Server çalıştıktan sonra ::8086 numaralı portta gelen istekleri dinlemeye başlayacaktır.
Kullanıcı adı "admin" şifre "admin123!"


