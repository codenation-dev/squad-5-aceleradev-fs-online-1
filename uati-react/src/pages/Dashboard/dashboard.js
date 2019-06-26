import React, { Component } from "react";
import ChartistGraph from "react-chartist";
import { Grid, Row, Col } from "react-bootstrap";

import api from '../../services/api';
import Card from '../../componentes/Card/Card';
import StatsCard from '../../componentes/StatsCard/StatsCard';

class Dashboard extends Component {
    constructor(props) {
        super(props)
        this.state = {
            alert: [],
            data: [],
            lineChartData: {
                labels: [],
                series: [
                  []
                ]
            },
            dataBar: {
                labels: [],
                series: [
                  []
                ]
            },
            lineChartData2: {
                labels: [],
                series: [
                  []
                ]
            },
            dataBar2: {
                labels: [],
                series: [
                  []
                ]
            }
        };
    }


    async componentDidMount() {
        const alerts = await api.get('dashboard/alerts');
                
        const numM = (a) => {
            const i = parseInt(a.month.substr(0,4))*100+parseInt(a.month.substr(5,2))
            return i
        }

        let baseData = alerts.data.data
        baseData.sort((a,b) => numM(a) - numM(b))

        const lineChartData = baseData.reduce((prev,curr) => {
            prev.labels.push(curr.month.substr(5,2)+"/"+curr.month.substr(0,4))
            prev.series[0].push(curr.alerts.clients.new_quantity)
            return prev
        },{
            labels: [],
            series: [[]]
        })
        const dataBar = baseData.reduce((prev,curr) => {
            prev.labels.push(curr.month.substr(5,2)+"/"+curr.month.substr(0,4))
            prev.series[0].push(curr.alerts.public_agent.customer_quantity)
            return prev
        },{
            labels: [],
            series: [[]]
        })

        const lineChartData2 = baseData.reduce((prev,curr) => {
            prev.labels.push(curr.month.substr(5,2)+"/"+curr.month.substr(0,4))
            prev.series[0].push(curr.alerts.bigger_salary.customer_quantity)
            return prev
        },{
            labels: [],
            series: [[]]
        })
        const dataBar2 = baseData.reduce((prev,curr) => {
            prev.labels.push(curr.month.substr(5,2)+"/"+curr.month.substr(0,4))
            prev.series[0].push(curr.alerts.bank_employee.customer_quantity)
            return prev
        },{
            labels: [],
            series: [[]]
        })

        this.setState({lineChartData,dataBar,lineChartData2,dataBar2})
        
    }    

    render() {        
          var lineChartOptions = {
            low: 0,
            showArea: true
          }
          var optionsBar = {
            low: 0,
            axisX: {
              labelInterpolationFnc: function(value, index) {
                return value
              }
            }
          };
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
                                            stats=""
                                            content={
                                                <div className="ct-chart">
                                                    <ChartistGraph
                                                        data={this.state.lineChartData}
                                                        type="Line"
                                                        options={lineChartOptions}
                                                    />
                                                </div>
                                            }
                                            
                                        />                                    
                            </Col>
                            <Col md={6}>
                                <Card
                                    id="chartActivity"
                                    title="Clientes que são funcionário público"
                                    category="Crescimento mensal"
                                    stats=""
                                    statsIcon="fa fa-check"
                                    content={
                                        <div className="ct-chart">
                                            <ChartistGraph
                                                data={this.state.dataBar}
                                                type="Bar"
                                                options={optionsBar}
                                            />
                                        </div>
                                    }
                                    
                                />
                            </Col>
                        </Row>
                        <Row>
                            <Col md={6}>               
                                        <Card
                                            statsIcon="fa fa-history"
                                            id="chartHours"
                                            title="Clientes com salário em destaque"
                                            category="Crescimento mensal"
                                            stats=""
                                            content={
                                                <div className="ct-chart">
                                                    <ChartistGraph
                                                        data={this.state.lineChartData2}
                                                        type="Bar"
                                                        options={lineChartOptions}
                                                    />
                                                </div>
                                            }
                                            
                                        />                                    
                            </Col>
                            <Col md={6}>
                                <Card
                                    id="chartActivity"
                                    title="Clientes que são funcionários do banco"
                                    category="Crescimento mensal"
                                    stats=""
                                    statsIcon="fa fa-check"
                                    content={
                                        <div className="ct-chart">
                                            <ChartistGraph
                                                data={this.state.dataBar2}
                                                type="Line"
                                                options={optionsBar}
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
