import React, { Component } from "react";
import axios from "axios";
import { Card, Header, Form, Input, Icon } from "semantic-ui-react";
import Button from "semantic-ui-react/dist/commonjs/elements/Button";

let endpoint = "http://localhost:8080";

class ToDoList extends Component {
    constructor(props) {
        super(props);

        this.state = {
            memberName: "",
            memberPhone: "",
            memberBirthday: new Date(),
            items: []
        };
    }

    componentDidMount() {
        this.getTask();
    }

    onChange = event => {
        this.setState({
            [event.target.name]: event.target.value
        });
    };

    onSubmit = () => {
        let { memberName, memberPhone, memberBirthday  } = this.state;
        // console.log("pRINTING memberName", this.state.memberName);
        if (memberName && memberPhone && memberBirthday)  {
            axios
                .post(
                    endpoint + "/api/member",
                    {
                        memberName,memberPhone,memberBirthday
                    },
                    {
                        headers: {
                            "Content-Type": "application/x-www-form-urlencoded"
                        }
                    }
                )
                .then(res => {
                    // this.getTask();
                    this.setState({
                        memberName: "",
                        memberPhone: "",
                        memberBirthday: new Date(),
                    });
                    console.log(res);
                });
        }
    };

    getTask = () => {
        axios.get(endpoint + "/api/member").then(res => {
            console.log(res);
            if (res.data) {
                this.setState({
                    items: res.data.map(item => {
                        let color = "yellow";

                        if (item.status) {
                            color = "green";
                        }
                        return (
                            <Card key={item._id} color={color} fluid>
                                <Card.Content>
                                    <Card.Header textAlign="left">
                                        <div style={{ wordWrap: "break-word" }}>{item.memberName}</div>
                                    </Card.Header>

                                    <Card.Meta textAlign="right">
                                        <Icon
                                            name="check circle"
                                            color="green"
                                            onClick={() => this.updateTask(item._id)}
                                        />
                                        <span style={{ paddingRight: 10 }}>Done</span>
                                        <Icon
                                            name="undo"
                                            color="yellow"
                                            onClick={() => this.undoTask(item._id)}
                                        />
                                        <span style={{ paddingRight: 10 }}>Undo</span>
                                        <Icon
                                            name="delete"
                                            color="red"
                                            onClick={() => this.deleteTask(item._id)}
                                        />
                                        <span style={{ paddingRight: 10 }}>Delete</span>
                                    </Card.Meta>
                                </Card.Content>
                            </Card>
                        );
                    })
                });
            } else {
                this.setState({
                    items: []
                });
            }
        });
    };

    updateTask = id => {
        axios
            .put(endpoint + "/api/memberName/" + id, {
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded"
                }
            })
            .then(res => {
                console.log(res);
                this.getTask();
            });
    };

    undoTask = id => {
        axios
            .put(endpoint + "/api/undoTask/" + id, {
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded"
                }
            })
            .then(res => {
                console.log(res);
                this.getTask();
            });
    };

    deleteTask = id => {
        axios
            .delete(endpoint + "/api/deleteTask/" + id, {
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded"
                }
            })
            .then(res => {
                console.log(res);
                this.getTask();
            });
    };
    render() {
        return (
            <div>
                <div className="row">
                    <Header className="header" as="h2">
                        PM
                    </Header>
                </div>
                <div className="row">
                    <Form onSubmit={this.onSubmit}>
                        <Input
                            type="text"
                            name="memberName"
                            onChange={this.onChange}
                            value={this.state.memberName}
                            fluid
                            placeholder="Name of member"
                        />
                        <Input
                            type="text"
                            name="memberPhone"
                            onChange={this.onChange}
                            value={this.state.memberPhone}
                            fluid
                            placeholder="Phone of member"
                        />
                        <Input
                            type="date"
                            name="memberBirthday"
                            onChange={this.onChange}
                            value={this.state.memberBirthday}
                            fluid
                            placeholder="Birthday of member"
                        />
                        { <Button >Create Task</Button> }
                    </Form>
                </div>
                <div className="row">
                    <Card.Group>{this.state.items}</Card.Group>
                </div>
            </div>
        );
    }
}

export default ToDoList;