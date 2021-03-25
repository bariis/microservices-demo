import React from 'react';
import 'bootstrap/dist/css/bootstrap.min.css'
import Navbar from 'react-bootstrap/Navbar';
import Nav from 'react-bootstrap/Nav';
import NavDropdown from 'react-bootstrap/NavDropdown';
import Jumbotron from 'react-bootstrap/Jumbotron';
import Button from 'react-bootstrap/Button';
import Image from 'react-bootstrap/Image';
import Media from 'react-bootstrap/Media';


const NavBar = () => {

    return (
        <Navbar collapseOnSelect expand="lg" bg="dark" variant="dark">
            <Navbar.Brand href="#home">Home</Navbar.Brand>
                <Nav>
                    <Button variant="secondary">Cart</Button>
                    <img src="../../public/cart.png" alt="Cart"/>
                </Nav>
        </Navbar>
    );
};

export default NavBar;