 IGI_APP/
   ├── igiInfra/                # dev ve prod clusterları için gerekli yapı
   │   ├── Chart.yaml           # Helm Chart Metadata
   │   ├── values.yaml          # Ortak environment ve port gibi global yapıların yönetimi, dev ve prod clusterları için gerekli constantları tutuyor 
   │   ├── values-dev.yaml      # Development ortamına özel değişkenler
   │   ├── values-prod.yaml     # Production ortamına özel değişkenler
   │   ├── templates/           # Kubernetes YAML dosyaları
   │   │   ├── namespace.yaml        # Namespace tanımı (dev/prod)
   │   │   ├── ingress.yaml          # Tüm API yönlendirme ve route yönetimi
   │   │   ├── storage.yaml          # Persistent Volume Yönetimi
   │   │   ├── configmap.yaml        # Ortak ConfigMap Template
   │   │   ├── secrets.yaml          # Ortak Secret Template
   │   │   ├── services/             # Tüm Mikroservislerin Deployment'ları ve Serviceleri
   │   │   │   ├── IGI_API-gateway.yaml  # (IGI_API-gateway podu)
   │   │   │   ├── IGI_API-people.yaml   
   │   │   │   ├── IGI_API-planet.yaml   
   │   │   │   ├── IGI_API-search.yaml   
   │   │   │   ├── IGI_CLIENT.yaml  
   ├── igiCli/                    # Sistem Yönetim Scriptleri
   │   ├── sys-setup.sh            # Setup scripti -- sys-diagnostic ve sys-config çağırıyor
   │   ├── sys-diag.sh             # Gerekli bağımlılıkları kontrol ve yükleme
   │   ├── sys-config.sh           #  githubdan public olan repoları çekerek igiAPI, IGICLIENT,IGIINFRA  içerisini dolduruyor Eğer gerekliyse helm k9s docker gibi teknolojileri projeye göre configure ediyor sistemi koşmadan ömnceki gerekli konfigürasyonları sağlıyor -- gerekliliği net değil
   ├── igiGateway/         # Kong Gatewayi sağlayan docker containerı. CORS, secıurtiy (auth değil), rate-limiting ve header gibi işlemleri yapan ve microsveris podlarına yönlendirme yapan docker containerı kendi içinde kong'un gerekliliklerini de ayağa kaldırır compose ile
   │   ├── ...
   │   ├── Dockerfile          
   ├── igiPeople/           # People mikroserviceini sağlayan docker containerı.  Kendi başına lokalde ayağa kalkıp postman ile test edilebilir ama client ve dışarıdan erişimi için gatewaye ihtiyaç var.
   │   ├── ...
   │   ├── Dockerfile            
   │   ├── go.mod               
   │   ├── go.sum              
   ├── igiPlanet/           # Planet mikroservisini sağlayan docker containerı
   │   ├── ...
   ├── igiSearch/           # Search mikroservisini sağlayan docker containerı.
   │   ├── ...
   ├── igiLib/              # Ortak Backend Kütüphanesini sağlayan static library. Containerı yok microservisler import edilerek kullanılıyor dependency olarak
   │   ├── helpers/              # ...Yardımcı Fonksiyonlar
   │   ├── middleware/           # ...Ortak middleware'ler (error handler vs.
   │   ├── config/               # ...Config ve Environment yönetiöi ( envler infra üzerinden dev veya prod diye yönetiliyor buradaki methodlar go servislerinin kodlamada bu değerleri kullanabilmelerine olanak sağlıyor - gerekliliği net değil)
   │   │   ├── ...
   ├── igiClient/                   # Frontend (React-Webpack server) - İleride belki micro-frontende geçebilir.
   │   ├── public/                  
   │   ├── src/                      
   │   ├── config/                   # Webpack config
   │   │   ├── config.dev.js        
   │   │   ├── config.prod.js      
   │   ├── docker/                   # Frontend için Docker yapılandırmaları
   │   │   ├── Dockerfile.dev        
   │   │   ├── Dockerfile.prod       
   │   ├── package.json              
   ├── README.md                     # Proje dokümantasyonu


# 1️⃣ Kubernetes Cluster (Okyanus) → Sanal bir yapıdır, node'ların iletişimini yönetir.
# 2️⃣ Node (Gemi) → Dış dünyaya açılan gerçek makinedir (internete bağlanır).
# 3️⃣ Pod (Gemi Bölmesi) → İçinde çalışan uygulamalar (container'lar) bulunur.
# 4️⃣ Container (Yük Kargosu) → Çalışan microservislerin kendisidir.
# 5️⃣ API Gateway (Liman-Gümrük Kapısı) → Dış dünyaya açılan tek noktadır.
# 6️⃣ ClusterIP (Gemi içi haberleşme ağı) → Pod’ların birbirleriyle iç iletişimi sağlar.
# 7️⃣ NodePort (Dış bağlantı kapısı) → Gerektiğinde dış dünyaya açılan kapıdır (Client gibi).

# ✅ Senin yapında doğru olan: Tüm trafik Gateway üzerinden yönetilmeli.
# ✅ Client ve backend servisleri ClusterIP ile haberleşmeli.
# ✅ Dış dünya sadece API Gateway'e erişebilmeli.

# 📌 Kubernetes Mimarisi: Okyanus ve Gemi Analoji ile Açıklama
# 1️⃣ Gateway ve Kullanıcı İletişimi
# Gateway, DNS’in yönlendirdiği ana gemidir ve tüm trafiği yönetir.
# Kullanıcı, igiapp.com adresine girdiğinde, istek önce Gateway’e ulaşır.
# Gateway, isteği Client pod’una yönlendirir.
# Client pod’u içinde bir container vardır, bu container React’ın statik build’ini barındırır.
# Kullanıcının tarayıcısı, bu statik dosyaları yükleyerek uygulamayı çalıştırır.
# 2️⃣ Backend ile İletişim
# Kullanıcı bir backend işlemi yaptığında (örneğin bir arama, form gönderimi), Client pod’u istek oluşturur.
# Bu istek önce Gateway’e gider.
# Gateway, ilgili microservice (MS) pod’una yönlendirme yapar.
# Hangi microservice gerekiyorsa, onun pod’una ulaşır ve isteği işler.
# Pod içindeki container, bu isteği karşılayan kodu çalıştırır.
# İşlem tamamlanınca, cevap Gateway’e döner.
# Gateway, cevabı Client pod’una iletir.
# Client, aldığı veriyi render ederek kullanıcıya gösterir.
# 3️⃣ Namespace ve Gemi Üretim Hatları
# Namespace, oluşturulan geminin nasıl bir ortamda çalışacağını belirler.
# Kubernetes okyanusuna açılacak bir gemi ile lokalimizde çalışan gemi birebir aynı değildir.
# Bu farkı yönetmek için iki farklı namespace (üretim hattı) kullanılır:
# Dev Namespace:
# Lokal geliştirme ortamında kullanılır.
# Tek bir node üzerinde çalışır (Minikube gibi).
# Prod ortamını taklit ederek, uygun portlarda çalışır.
# Prod Namespace:
# Gerçek Kubernetes okyanusunda çalışacak versiyondur.
# Okyanusa çıkmaya hazır, optimize edilmiş stabil bir gemi üretir.
# Bu gemi, Dev ortamında test edilmiş ve stabil hale getirilmiş imajlardan oluşur.
# 4️⃣ Dev’den Prod’a Geçiş ve Image Mantığı
# Prod’a çıkacak gemi, Dev ortamında üretilmiş stabil bir versiyonun binary halidir.
# Bu yüzden, prod versiyonu bir "image" olarak oluşturulur.
# Image, bir geminin tüm çalışma mantığını içeren, her yerde çalışabilen bir pakettir.
# Bu image, Kubernetes tarafından uygun node’lara (gemilere) dağıtılarak çalıştırılır.
# 📌 Özet: Profesyonel Kubernetes Akışı
# 1️⃣ Kullanıcı, Gateway’e ulaşır.
# 2️⃣ Gateway, isteği Client pod’una iletir.
# 3️⃣ Client pod’u UI’ı yükler ve kullanıcıya sunar.
# 4️⃣ Backend işlemleri gerektiğinde, Client pod’u Gateway’e istek atar.
# 5️⃣ Gateway, ilgili microservice pod’una isteği yönlendirir.
# 6️⃣ Microservice pod’u işlemi yapar, Gateway’e cevap döner.
# 7️⃣ Gateway, cevabı Client pod’una iletir.
# 8️⃣ Client, kullanıcıya veriyi gösterir.
# 9️⃣ Dev ortamı ve Prod ortamı namespace’ler ile ayrılır.
# 🔟 Prod ortamına çıkacak gemi, Dev ortamında test edilmiş stabil bir image olur.

# Bu, senin mimarine uygun olarak profesyonel projelerde nasıl çalıştığını net bir şekilde anlatır.