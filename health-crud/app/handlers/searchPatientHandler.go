package handlers

import (
	"fmt"
	"health-crud/app/utils"
	"net/http"
)

const searchResponseTemplate = `
<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Resultado da Busca</title>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600&display=swap" rel="stylesheet">
    <style>
        body { font-family: 'Inter', sans-serif; background-color: #f8fafc; color: #1e293b; display: flex; justify-content: center; align-items: center; min-height: 100vh; margin: 0; padding: 2rem;}
        .container { background: white; padding: 3rem; border-radius: 12px; box-shadow: 0 10px 25px rgba(0,0,0,0.05); text-align: center; max-width: 600px; width: 100%%; }
        h1 { margin-top: 0; color: %s; }
        p { color: #64748b; margin-bottom: 2rem; line-height: 1.5; }
        .details { text-align: left; background: #f1f5f9; padding: 1.5rem; border-radius: 8px; margin-bottom: 2rem; }
        .details p { margin: 0.5rem 0; color: #334155; }
        .details strong { color: #0f172a; }
        .btn { display: inline-block; background-color: #3b82f6; color: white; padding: 0.8rem 1.5rem; border-radius: 8px; text-decoration: none; font-weight: 600; transition: all 0.2s; }
        .btn:hover { background-color: #2563eb; transform: translateY(-2px); }
    </style>
</head>
<body>
    <div class="container">
        <h1>%s</h1>
        %s
        <a href="/" class="btn">Voltar ao Início</a>
    </div>
</body>
</html>
`

func SearchPatientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	cpf := r.FormValue("cpf")
	patient, err := utils.GetPatientByCPF(cpf)
	if err != nil {
		msg := fmt.Sprintf("<p>%v</p>", err)
		fmt.Fprintf(w, searchResponseTemplate, "#ef4444", "Erro na Busca", msg)
		return
	}

	if patient == nil {
		msg := fmt.Sprintf("<p>Nenhum paciente encontrado com o CPF: <strong>%s</strong></p>", cpf)
		fmt.Fprintf(w, searchResponseTemplate, "#f59e0b", "Paciente Não Encontrado", msg)
		return
	}

	details := fmt.Sprintf(`
		<div class="details">
			<p><strong>Nome:</strong> %s</p>
			<p><strong>CPF:</strong> %s</p>
			<p><strong>Data de Nascimento:</strong> %s</p>
			<p><strong>Tipo Sanguíneo:</strong> %s</p>
			<p><strong>Telefone:</strong> %s</p>
			<p><strong>E-mail:</strong> %s</p>
			<p><strong>Anotações Médicas:</strong> %s</p>
		</div>
	`, patient.Name, patient.CPF, patient.BirthDate, patient.BloodType, patient.Phone, patient.Email, patient.MedicalNotes)

	fmt.Fprintf(w, searchResponseTemplate, "#3b82f6", "Detalhes do Paciente", details)
}
