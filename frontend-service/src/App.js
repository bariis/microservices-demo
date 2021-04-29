import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css'

import { BrowserRouter, Route, Switch } from 'react-router-dom';

import Navbar from "./components/NavBar"
import Cart from './components/Cart';

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Navbar />
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