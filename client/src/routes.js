import React from 'react'
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import About from './pages/About';
import Register from './pages/Register';
import ListPlanets from './pages/ListPlanets';



export default function Routes() {
    return(
        <BrowserRouter>
            <Switch>
                <Route path="/" exact component={About} />
                <Route path="/register" component={Register} />
                <Route path="/planets" component={ListPlanets} />
            </Switch>
        </BrowserRouter>
    );
}
