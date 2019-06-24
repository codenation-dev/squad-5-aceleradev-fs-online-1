import React, { Component } from "react";
import ChartistGraph from "react-chartist";
import { Grid, Row, Col } from "react-bootstrap";

import api from '../../services/api';
import Card from '../../componentes/Card/Card';
import StatsCard from '../../componentes/StatsCard/StatsCard';
import {
    dataSales,
    optionsSales,
    responsiveSales,    
    dataBar,
    optionsBar,
    responsiveBar,   
} from '../../variables/Variables';

class Dashboard extends Component {
    state = {
            alert: [],
        };

    async componentDidMount() {
        const response = await api.get('dashboard/alerts');

        this.setState({alert: response.data})
                
        let details = [];  
                          
        for (var i in response.data) {
            details.push({ name: i, value: response.data[i] })
            //console.log(details)            
        }       
        this.setState({ data:  details })

        console.log(this.state.data[0])
        
    }    

    //{this.state.data.map((alert) => (
    //))}

    render() {        
        return (                          
                <div className="content">
                    <Grid fluid>
                        <Row>      
                        {this.state.data.map(alert => (                      
                        <Col lg={2} sm={4}>                            
                                <StatsCard                                    
                                    statsText={alert.name}
                                    statsValue={alert.value}
                                    statsIcon={<i className="fa fa-calendar-o" />}
                                    statsIconText="Last day"
                                />                                
                            </Col>   
                            ))}                         
                        </Row>
                        <br />
                        <Row>
                            <Col md={6}>               
                                        <Card
                                            statsIcon="fa fa-history"
                                            id="chartHours"
                                            title="Clientes cadastrados"
                                            category="Crescimento mensal"
                                            stats="Informações verificadas"
                                            content={
                                                <div className="ct-chart">
                                                    <ChartistGraph
                                                        data={dataSales}
                                                        type="Line"
                                                        options={optionsSales}
                                                        responsiveOptions={responsiveSales}
                                                    />
                                                </div>
                                            }
                                            
                                        />                                    
                                }
                                )}
                            </Col>
                            <Col md={6}>
                                <Card
                                    id="chartActivity"
                                    title="Crescimento"
                                    category="Clientes cadastrados"
                                    stats="Informações verificadas"
                                    statsIcon="fa fa-check"
                                    content={
                                        <div className="ct-chart">
                                            <ChartistGraph
                                                data={dataBar}
                                                type="Bar"
                                                options={optionsBar}
                                                responsiveOptions={responsiveBar}
                                            />
                                        </div>
                                    }
                                    
                                />
                            </Col>
                        </Row>
                        
                    </Grid>
                </div>  
        );   
}}

export default Dashboard;
