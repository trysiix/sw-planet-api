import React, { useState } from 'react';
import { Link, useHistory } from 'react-router-dom';
import api from '../../services/api';

import './styles.css';
import img from '../../assets/dstar.png';
import { FiArrowLeft } from 'react-icons/fi';


export default function NewIncident() {
    const [name, setName] = useState('');
    const [weather, setWeather] = useState('');
    const [terrain, setTerrain] = useState('');
    
    const history = useHistory();


    async function handleNewRegister (submitData) {
        submitData.preventDefault();

        const data = {
            name,
            weather,
            terrain,
        }

          console.log(data);
        try {
            await api.post('api/planet', data)
            history.push('/planets');
        } catch (error) {
            alert('Erro ao cadastrar planeta, tente novamente.');
        } 
    }

    return(
        <div className="new-planet">
            <div className="content">
                <section>
                    <form>
                        <img src={img} alt="planet"/>
                        <h1>Cadastrar Planeta</h1>
                        <p>Insira os dados do planeta.</p>

                        <Link className="back-link" to="/planets">
                            <FiArrowLeft size={16} color="#E02041"/>
                             Listar Planetas
                        </Link>
                    </form>
                </section>
                <form onSubmit={handleNewRegister}>
                    <input 
                        placeholder="Nome" 
                        value={name}
                        onChange={submitData => setName(submitData.target.value)}
                    />
                     <input 
                        placeholder="Clima" 
                        value={weather}
                        onChange={submitData => setWeather(submitData.target.value)}
                    />
                    <input 
                        placeholder="Terreno" 
                        value={terrain}
                        onChange={submitData => setTerrain(submitData.target.value)}
                    />

                    <button className="button" type="submit">Cadastrar</button>
                </form>
            </div>
        </div>
    );
}
