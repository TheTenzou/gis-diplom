import MySidebar from "../mapPage/Sidebar";
import decodeToken from "../adminPage/decodeJwt";
import userData from "../adminPage/AdminPanel";

import React, { Component } from "react";
import { Navbar, Nav, Button, Form, Collapse, BImg } from "bootstrap-4-react";
import Sidebar from "react-sidebar";

export default class MainNavigation extends Component {
  constructor(props) {
    super(props);
    this.state = {
      sidebarOpen: false,
    };
    this.onSetSidebarOpen = this.onSetSidebarOpen.bind(this);
  }

  onSetSidebarOpen(e, open) {
    if (open) {
      e.preventDefault();
    }
    this.setState({ sidebarOpen: open });
  }

  render() {
    const bootstrap_icon =
      "https://static.tildacdn.com/tild3236-6566-4134-b464-623730343364/square-md-a.png";
    return (
      <React.Fragment>
        {this.props.sideBar && (
          <Sidebar
            sidebar={<MySidebar />}
            open={this.state.sidebarOpen}
            pullRight={true}
            onSetOpen={this.onSetSidebarOpen}
            styles={{
              sidebar: { background: "white", zIndex: "1001", width: "300px" },
            }}
          ></Sidebar>
        )}

        <Navbar expand="lg" dark bg="dark" mb="3">
          <Navbar.Brand href="/">
            <BImg
              src={bootstrap_icon}
              width="30"
              height="30"
              display="inline-block"
              align="top"
              mr="1"
            />
          </Navbar.Brand>
          <Navbar.Toggler target="#navbar" />
          <Collapse navbar id="navbar">
            <Navbar.Nav mr="auto">
              <Nav.ItemLink href="/" active>
                Главная {this.props.sideBar}
              </Nav.ItemLink>
              <Nav.ItemLink href="/tsoddSchedules/1">
                План обслуживания ТСОДД
              </Nav.ItemLink>
              <Nav.ItemLink href="/table2">
                План по модернизации УДС
              </Nav.ItemLink>
              <Nav.ItemLink href="/about">
                Дополнительная информация
              </Nav.ItemLink>
              <div id="modalRoot"></div>
            </Navbar.Nav>
            <Form inline my="2 lg-0">
              <ValidityTest />
              {this.props.sideBar && (
                <Button
                  outline
                  light
                  mr="sm-2"
                  my="2 sm-2"
                  onClick={(e) => this.onSetSidebarOpen(e, true)}
                >
                  Сортировки
                </Button>
              )}
              <Button outline info my="2 sm-0" as="a" href="/auth">
                Войти
              </Button>
            </Form>
          </Collapse>
        </Navbar>
      </React.Fragment>
    );
  }
}

function ValidityTest() {
  try {
    var jwt = decodeToken(localStorage.getItem("accessToken"));

    if (jwt["roles"].indexOf("ROLE_ADMIN") !== -1) {
      return (
        <Button outline warning mr="sm-2" as="a" href="/adminPanel">
          Админская панель
        </Button>
      );
    } else {
      return <p>{console.log("You're not an admin")}</p>;
    }
  } catch {}

  try {
    console.log(userData);
  } catch {}

  return <p>{console.log("You're not an admin")}</p>;
}
