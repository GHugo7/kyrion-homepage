import { useEffect, useState } from "react"
import './index.css'
import logo from './assets/logo.png'    

interface Service {
    titre: string
    description: string
    categories: string
    url: string
    online: boolean
}

interface Categorie {
    nom: string
    services: Service[]
}

export default function App() {
    const [ categorie, setCategorie ] = useState<Categorie[]>([]);
    const [ selectedService, setSelectedService] = useState<Service | null>(null);

    const api = "http://192.168.1.69:1818/api/services"
    useEffect(() => { 
        fetch(api)
        .then(reponse => reponse.json())
        .then(data => setCategorie(data))
    }, [])
    
    if (categorie.length === 0) {
        return (
            <p className="loading">Chargement...</p>
        )
    }

    return (
        <>
        <header className="site-header">
            <img src={logo} alt="Kyrion" className="logo" />
            <button className="btn-admin">Admin</button>
        </header>
        <h1>Kyrion</h1>
        {categorie.map((cat) => (
            <div className="categorie-box" key={cat.nom}>
                <h2> {cat.nom} </h2>
                <div className="services-grid">
                    {cat.services.map((s) => (
                    <div className="service-box" key={s.titre} onClick={() => setSelectedService(s)}>
                        <h3> 
                            <span className={s.online ? "status-dot online" : "status-dot offline"}></span>
                            {s.titre} 
                        </h3>
                        <p> {s.description} </p>
                        {s.url && (
                            <a href={s.url} className="btn-aller" target="_blank" rel="noopener noreferrer" onClick={(e) => e.stopPropagation()}>Aller</a>
                        )}
                    </div>
                ))}
                </div>
            </div>
        ))}
        {selectedService && (
            <div className="modal-overlay" onClick={() => setSelectedService(null)}>
                <div className="modal-content" onClick={(e) => e.stopPropagation()}>
                    <h2>{selectedService.titre}</h2>
                    {/* Details */}
                    
                </div>
            </div>
        )}
        </>
    )
}