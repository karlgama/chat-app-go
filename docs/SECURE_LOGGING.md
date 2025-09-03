# Sistema de Logging Seguro com Trace ID

## 🔐 **Proteção de Dados PII**

Este sistema implementa logging seguro que automaticamente mascara dados sensíveis em ambientes de produção, mantendo apenas o `trace_id` para rastreabilidade.

## 🎯 **Funcionalidades**

### 1. **Trace ID Automático**
- Extrai `X-Trace-ID` do header ou gera UUID automaticamente
- Adiciona o trace ID em todas as respostas
- Disponível no contexto para todos os controllers

### 2. **Mascaramento Automático de PII**
- **Ambiente Local + Debug**: Dados completos são logados
- **Outros ambientes**: Dados sensíveis são mascarados

### 3. **Campos Sensíveis Identificados**
```go
// Campos que são automaticamente mascarados:
"email", "password", "token", "jwt", "secret",
"phone", "cpf", "cnpj", "credit_card", "ssn",
"name", "full_name", "address", "user_id"
```

## 🚀 **Exemplos de Uso**

### Logging Seguro no Controller:
```go
// ✅ Forma correta - dados sensíveis são mascarados automaticamente
utils.SafeLogInfo(c, "Login attempt", map[string]interface{}{
    "email": "user@example.com",  // Será mascarado em produção
    "action": "login",            // Sempre visível
})
```

### Output por Ambiente:

**Local + Debug:**
```json
{
  "trace_id": "123e4567-e89b-12d3-a456-426614174000",
  "ip": "127.0.0.1",
  "endpoint": "POST /login",
  "email": "user@example.com",
  "action": "login",
  "level": "info",
  "msg": "Login attempt"
}
```

**Produção:**
```json
{
  "trace_id": "123e4567-e89b-12d3-a456-426614174000", 
  "ip": "127.0.0.1",
  "endpoint": "POST /login",
  "email_masked": "u***@example.com",
  "action": "login",
  "level": "info",
  "msg": "Login attempt"
}
```

## 🛠 **APIs Disponíveis**

```go
// Logs com diferentes níveis
utils.SafeLogInfo(c, "message", fields)
utils.SafeLogWarn(c, "message", fields) 
utils.SafeLogError(c, "message", fields)

// Extrair trace ID manualmente
traceID := middlewares.GetTraceID(c)

// Criar campos seguros manualmente
safeFields := utils.SafeLogFields(c, fields)
```

## 🔍 **Rastreabilidade**

Todas as requisições agora têm um trace ID único que permite:
- Rastrear uma requisição específica nos logs
- Correlacionar logs entre diferentes serviços
- Debugging sem expor dados sensíveis

## ⚙️ **Configuração**

O mascaramento é ativado automaticamente baseado em:
- **Ambiente**: Deve ser diferente de "local"  
- **Log Level**: Deve ser diferente de "debug"

Para debug completo:
```env
APP_ENV=local
LOG_LEVEL=debug
```
