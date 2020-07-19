import React , { useState }  from 'react';
import { Link , useHistory} from 'react-router-dom';
import api from '../../services/api';


import './style.css' ;
import logoImg from '../../assets/sw_name.png';
import { FiLogIn } from 'react-icons/fi';


export default function About() {
    return (
        <div className="home-container">
            <section className="form">
                <img className="header-icon" src= { logoImg } alt="Star Wars"/>

                <form>
                    <h1><b>Welcome to the dark side</b></h1>
                    <h2><b>We have planets</b></h2>
                    <p><b>Features: </b></p>
                    <p><b> <i>Cadastrar Planeta</i></b></p>
                    <p><b> <i>Listar Planetas</i></b></p>
                    <p><b> <i> Buscar planeta por id</i></b></p>
                    <p><b> <i> Buscar planetas por filtro</i></b></p>

                    <span><i>Criado por Carlos Henrique</i></span>

                <Link className="back-link" to="/register">
                   <FiLogIn size={16} color="#E02041"/>
                   Cadastrar Planeta
                </Link>
                <Link className="back-link" to="/planets">
                   <FiLogIn size={16} color="#E02041"/>
                   Listar Planeta
                </Link>
                </form>
            </section>
        </div>
    );
}
