import React, { useState , useEffect } from 'react';
import { Link } from 'react-router-dom';
import api from '../../services/api';

import logoImg from '../../assets/sw_name.png';
import { FiTrash2 } from 'react-icons/fi';
import './styles.css';

export default function ListPlanets() {
    const [planets, setPlanets] = useState([]);

     useEffect(() => {
         api.get('api/planet').then(response => {
            setPlanets(response.data);
        })
    });

    async function handleDeletePlanet(id) {
        try {
            await api.delete(`api/planet/del/${id}`);

            setPlanets(planets.filter(planet => planet.id !== id));
        } catch (error) {
            alert('Erro ao deletar planeta, tente novamente.');
        }
    }

    return(
       <div className="planet-container">
            <header>
                <img src={logoImg} alt="sw"/>

                <Link className="button" to="/register">Cadastrar planeta</Link>
                <Link className="button" to="/">Sobre</Link>
            </header>

            <h1 align="center">Planetas</h1>

            <ul>
                {planets?planets.map(planet => (
                <li key={planet._id}>
                        <strong>
                            Planeta: 
                            <i>{planet.name}</i>
                        </strong>
                        

                        <strong>
                            Clima: 
                            <i>{planet.weather}</i>
                        </strong>
                        

                        <strong>
                            Tipo do Terreno: 
                            <i>{planet.terrain}</i>
                        </strong>
                        

                        <strong>
                            Quantidade de aparições em filmes: 
                            <i>{planet.numberofappearances}</i>
                        </strong>

                        <button onClick={() => handleDeletePlanet(planet._id)}>
                            <FiTrash2 size={20} color="#a8a8b3"/>
                        </button>
                </li>
                )):
                <li key="No-index">
                        <strong>Nenhum planeta foi encontrado</strong>
                </li>}
            </ul>
       </div>
    );
}
