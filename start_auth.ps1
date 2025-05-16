# Set environment variables
$env:JWT_SECRET = "virtual_decanat_secret_key_for_development"
$env:JWT_EXPIRATION = "24h"
$env:PORT = "8081"

# Build the authentication service
Write-Host "Building authentication service..."
Set-Location -Path auth-service/backend
go build -o auth-service.exe cmd/main.go

# Run the authentication service
Write-Host "Starting authentication service on port $env:PORT..."
Start-Process -FilePath ".\auth-service.exe" -NoNewWindow

Write-Host "Authentication service started. Use CTRL+C to stop."
Write-Host
Write-Host "API Documentation:"
Write-Host "1. Register: POST http://localhost:8081/register"
Write-Host "   Body: { \"username\": \"user\", \"email\": \"user@example.com\", \"password\": \"password\", \"first_name\": \"First\", \"last_name\": \"Last\", \"role\": \"student\" }"
Write-Host
Write-Host "2. Login: POST http://localhost:8081/login"
Write-Host "   Body: { \"username\": \"user\", \"password\": \"password\" }"
Write-Host
Write-Host "3. Get Profile: GET http://localhost:8081/profile"
Write-Host "   Headers: Authorization: Bearer {token}"
Write-Host
Write-Host "4. Update Profile: PUT http://localhost:8081/profile"
Write-Host "   Headers: Authorization: Bearer {token}"
Write-Host "   Body: { \"first_name\": \"Updated\", \"last_name\": \"Name\" }"
Write-Host
Write-Host "5. Change Password: POST http://localhost:8081/change-password"
Write-Host "   Headers: Authorization: Bearer {token}"
Write-Host "   Body: { \"old_password\": \"password\", \"new_password\": \"newpassword\" }"

# Wait for user input
$host.UI.RawUI.ReadKey("NoEcho,IncludeKeyDown") | Out-Null
Write-Host "Stopping..."

# Return to root directory
Set-Location -Path ../.. 