import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css'

import { BrowserRouter, Route, Switch } from 'react-router-dom';

import Navbar from "./components/NavBar"
import Cart from './components/Cart';
import Catalog from './components/Catalog';

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Navbar />
        <Catalog />
        <Switch>
        <Route path="/cart">
          <Cart />
        </Route>
        </Switch>  
      </BrowserRouter>
    </div>
  );
}

export default App;