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
    <link href="https://fonts.googleapis.com/css2?family=Outfit:wght@400;600;700&display=swap" rel="stylesheet">
    <style>
        :root {
            color-scheme: light;
        }
        * {
            box-sizing: border-box;
        }
        body {
            margin: 0;
            min-height: 100vh;
            display: flex;
            flex-direction: column;
            background-color: #EBF2F8;
            background-image: radial-gradient(circle, rgba(27,95,140,0.07) 1px, transparent 1px);
            background-size: 26px 26px;
            font-family: 'Outfit', sans-serif;
            color: #1A2C3A;
            padding-top: 64px;
        }
        header {
            position: fixed;
            top: 0;
            left: 0;
            right: 0;
            height: 64px;
            background: #1B5F8C;
            display: flex;
            align-items: center;
            justify-content: space-between;
            padding: 0 1.5rem;
            z-index: 20;
        }
        .brand {
            display: flex;
            align-items: center;
            gap: 0.85rem;
        }
        .brand svg {
            display: block;
        }
        .brand-title {
            color: #FFFFFF;
            font-weight: 700;
            font-size: 1rem;
            margin: 0;
        }
        .brand-subtitle {
            color: rgba(255, 255, 255, 0.65);
            font-size: 0.92rem;
            margin: 0;
        }
        main {
            flex: 1;
            display: flex;
            justify-content: center;
            align-items: center;
            padding: 2rem;
        }
        .card {
            width: 100%%;
            max-width: 480px;
            background: #FFFFFF;
            border-radius: 14px;
            padding: 2.5rem;
            box-shadow: 0 4px 20px rgba(27,95,140,0.08);
            border-top: 4px solid var(--action-color, #1B5F8C);
            animation: slideUp 0.38s ease both;
            position: relative;
        }
        .card[data-color="#3b82f6"] { --action-color: #1B5F8C; }
        .card[data-color="#f59e0b"] { --action-color: #1B5F8C; }
        .card[data-color="#ef4444"] { --action-color: #C0392B; }
        .card .icon::before {
            content: "🔍";
        }
        .card[data-color="#f59e0b"] .icon::before {
            content: "🔎";
        }
        .card[data-color="#ef4444"] .icon::before {
            content: "⚠️";
        }
        .icon {
            display: block;
            font-size: 2.5rem;
            text-align: center;
            margin-bottom: 1.25rem;
        }
        h1 {
            font-size: 1.4rem;
            font-weight: 700;
            text-align: center;
            margin: 0 0 1rem;
            color: #1A2C3A;
        }
        .message {
            font-size: 0.9rem;
            color: #5D7889;
            line-height: 1.6;
            text-align: center;
            margin: 0 0 1.75rem;
        }
        .details {
            background: #F6FAFD;
            border: 1px solid #C4D9E8;
            border-radius: 10px;
            padding: 1.25rem;
            text-align: left;
            margin: 0 0 1.75rem;
        }
        .details p {
            margin: 0.6rem 0;
            font-size: 0.9rem;
            color: #1A2C3A;
        }
        .details p strong {
            font-weight: 700;
        }
        .button {
            display: inline-flex;
            align-items: center;
            justify-content: center;
            gap: 0.35rem;
            background: #1B5F8C;
            color: #FFFFFF;
            border: none;
            border-radius: 9px;
            padding: 11px 28px;
            text-decoration: none;
            font-weight: 600;
            transition: background 0.18s ease, transform 0.15s ease;
        }
        .button:hover {
            background: #134567;
            transform: translateY(-1px);
        }
        footer {
            width: 100%%;
            text-align: center;
            padding: 1rem 0;
            color: rgba(255,255,255,0.5);
            background: #134567;
            font-size: 0.78rem;
        }
        @keyframes slideUp {
            from {
                opacity: 0;
                transform: translateY(16px);
            }
            to {
                opacity: 1;
                transform: translateY(0);
            }
        }
    </style>
</head>
<body>
    <header>
        <div class="brand">
            <svg width="26" height="26" viewBox="0 0 28 28" fill="none" xmlns="http://www.w3.org/2000/svg">
                <rect x="10" y="2" width="8" height="24" rx="2.5" fill="white"/>
                <rect x="2" y="10" width="24" height="8" rx="2.5" fill="white"/>
            </svg>
            <div>
                <p class="brand-title">HealthCare CRUD</p>
                <p class="brand-subtitle">Sistema de Gestão de Pacientes</p>
            </div>
        </div>
        <p class="brand-subtitle">Sistema de Gestão de Pacientes</p>
    </header>
    <main>
        <div class="card" data-color="%s">
            <span class="icon"></span>
            <h1>%s</h1>
            <div class="message">%s</div>
            <a href="/" class="button">Voltar ao Início</a>
        </div>
    </main>
    <footer>© 2025 HealthCare CRUD · Projeto Acadêmico</footer>
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
