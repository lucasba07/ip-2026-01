package handlers

import (
	"fmt"
	"health-crud/app/utils"
	"net/http"
)

const deleteResponseTemplate = `
<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Resultado</title>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600&display=swap" rel="stylesheet">
    <style>
        body { font-family: 'Inter', sans-serif; background-color: #f8fafc; color: #1e293b; display: flex; justify-content: center; align-items: center; min-height: 100vh; margin: 0; }
        .container { background: white; padding: 3rem; border-radius: 12px; box-shadow: 0 10px 25px rgba(0,0,0,0.05); text-align: center; max-width: 500px; width: 90%%; }
        h1 { margin-top: 0; color: %s; }
        p { color: #64748b; margin-bottom: 2rem; line-height: 1.5; }
        .btn { display: inline-block; background-color: #f1f5f9; color: #475569; padding: 0.8rem 1.5rem; border-radius: 8px; text-decoration: none; font-weight: 600; transition: all 0.2s; }
        .btn:hover { background-color: #e2e8f0; transform: translateY(-2px); }
    </style>
</head>
<body>
    <div class="container">
        <h1>%s</h1>
        <p>%s</p>
        <a href="/" class="btn">Voltar ao Início</a>
    </div>
</body>
</html>
`

func DeletePatientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	cpf := r.FormValue("cpf")

	err := utils.DeletePatient(cpf)
	if err != nil {
		fmt.Fprintf(w, deleteResponseTemplate, "#ef4444", "Erro ao excluir", err.Error())
		return
	}

	fmt.Fprintf(w, deleteResponseTemplate, "#10b981", "Excluído!", "O registro do paciente foi excluído permanentemente.")
}
