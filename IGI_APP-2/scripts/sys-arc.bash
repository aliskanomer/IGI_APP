IGI_APP/
  ├── IGI_INFRA/                 # Kubernetes & Deployment Yönetimi (Helm Kullanılıyor)
  │   ├── helm/                   # Helm Chart Yapısı
  │   │   ├── Chart.yaml           # Helm Chart Metadata
  │   │   ├── values.yaml          # Ortak environment değişkenleri (global)
  │   │   ├── values-dev.yaml      # Development ortamına özel değişkenler
  │   │   ├── values-prod.yaml     # Production ortamına özel değişkenler
  │   │   ├── templates/           # Kubernetes YAML dosyaları
  │   │   │   ├── namespace.yaml        # Namespace tanımı (dev/prod)
  │   │   │   ├── ingress.yaml          # Tüm API yönlendirme ve route yönetimi
  │   │   │   ├── kong-deployment.yaml  # Kong Gateway için Kubernetes Deployment (IGI_API-gateway podu)
  │   │   │   ├── kong-service.yaml     # Kong Gateway için Kubernetes Service
  │   │   │   ├── storage.yaml          # Persistent Volume Yönetimi
  │   │   │   ├── redis.yaml            # Redis Deployment & Service
  │   │   │   ├── configmap.yaml        # Ortak ConfigMap Template
  │   │   │   ├── secrets.yaml          # Ortak Secret Template
  │   │   │   ├── services/             # Tüm Mikroservislerin Deployment'ları
  │   │   │   │   ├── IGI_API-gateway.yaml 
  │   │   │   │   ├── IGI_API-people.yaml   
  │   │   │   │   ├── IGI_API-planet.yaml   
  │   │   │   │   ├── IGI_API-search.yaml   
  │   │   │   │   ├── IGI_CLIENT.yaml  
  │   │   │   ├── monitoring/            # Monitoring & Loglama Deployment'ları
  │   │   │   │   ├── grafana.yaml       
  │   │   │   │   ├── loki.yaml        
  │   │   │   │   ├── promtail.yaml    
  ├── scripts/                    # Sistem Yönetim Scriptleri
  │   ├── sys-diag.sh             # Gerekli bağımlılıkları kontrol ve yükleme
  │   ├── sys-setup.sh            # Setup scripti -- sys-diagnostic ve sys-config çağırıyor
  ├── IGI_API/                     # Backend (Microservices & API Gateway)
  │   ├── IGI_API-gateway/         # Kong Gateway 
  │   │   ├── ...
  │   │   ├── Dockerfile          
  │   ├── IGI_API-people/           # People Servisi Microservice
  │   │   ├── ...
  │   │   ├── Dockerfile            
  │   │   ├── go.mod               
  │   │   ├── go.sum              
  │   ├── IGI_API-planet/           # Planet Servisi
  │   │   ├── ...
  │   ├── IGI_API-search/           # Search Servisi
  │   │   ├── ...
  │   ├── IGI_API-lib/              # Ortak Backend Kütüphanesi
  │   │   ├── helpers/              # Yardımcı Fonksiyonlar
  │   │   ├── middleware/           # Ortak middleware'ler (error handler vs.
  │   │   ├── config/               # Config ve Environment yönetimi
  │   │   ├── ...
  ├── IGI_CLIENT/                   # Frontend 
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


  # ----------------------------------------
# 🔹 IGI_APP Mimarisi Çalışma Mantığı: (EKSIK GIUNCELLLE)
#
# 1️⃣ **IGI_INFRA (Kubernetes & Helm)**
#    - Kubernetes cluster yönetimi için Helm kullanılır.
#    - Mikroservisler, Kong Gateway ve Redis gibi bağımlılıklar burada tanımlanır.
#    - `helm/templates/` içinde `Deployment`, `Service` ve `ConfigMap` gibi YAML dosyaları bulunur.
#
# 2️⃣ **IGI_API (Backend)**
#    - Mikroservisler burada yer alır (People, Planet, Search).
#    - Kong Gateway (`IGI_API-gateway`) burada çalışır ve tüm API isteklerini yönlendirir.
#    - Her servis bağımsız olarak çalışır ve Kong Gateway üzerinden erişilir.
#
# 3️⃣ **IGI_CLIENT (Frontend)**
#    - React tabanlı frontend, API isteklerini `IGI_API-gateway` üzerinden yapar.
#    - Webpack ile derlenir, Docker içinde çalıştırılır.
#
# 4️⃣ **scripts (Otomasyon)**
#    - `sys-diag.sh`: Gerekli bağımlılıkları kontrol eder ve yükler.
#    - `sys-setup.sh`: Tüm kurulumu ve yapılandırmayı başlatır.
#
# 5️⃣ **Çalıştırma Adımları**
#    - `setup` komutu tüm süreci başlatır (`sys-diag.sh` ve `sys-setup.sh` çalışır).
#    - Kubernetes deployment işlemleri Helm üzerinden yapılır.
#    - Mikroservisler `kubectl apply -f IGI_API-xxx.yaml` komutlarıyla ayağa kaldırılır.
#    - Kong Gateway (`IGI_API-gateway`) API yönetimini sağlar.
#    - IGI_CLIENT, backend servislerine `IGI_API-gateway` üzerinden erişir.
#
# ----------------------------------------

