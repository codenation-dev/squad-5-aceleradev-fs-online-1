import React, { Component } from 'react'
import {withRouter} from 'react-router-dom';
import Link from '../../componentes/Link/Link'
import Botao from '../../componentes/Botao/Botao'
import Legenda from '../../componentes/Legenda/Legenda'
import Campo from '../../componentes/Campo/Campo'
import './Login.css'
import api from '../../services/api'

class Login extends Component {
  constructor(props) {
    super(props)

    this.emailRef = React.createRef() // { current: null }
    this.senhaRef = React.createRef()
    
    this.state = { desabilitado: true }
  }

  enviaDados = async (evento) => {
    evento.preventDefault()

    const campoEmail = this.emailRef.current
    const campoSenha = this.senhaRef.current

    const dados = {
      username: campoEmail.getValor(),
      password: campoSenha.getValor()
    }

    try {
      await api.post("auth", dados)
      .then(res => localStorage.setItem('TOKEN', res.data.token))
    } catch (e) {
      alert("Usuário e senha inválido")
      return
    }

    this.props.history.push('/')

  }

  habilitaOuDesabilitaBotao = () => {
    const campoEmail = this.emailRef.current
    const campoSenha = this.senhaRef.current

    if (campoEmail.temErro() || campoSenha.temErro()) {
      this.setState({ desabilitado: true })
    } else {
      this.setState({ desabilitado: false })
    }
  }

  render() {
    return (
      <main className="login">
        <div className="form">      
        <h1>Login</h1>
        {/* <p>Entre com seu usuário e senha.</p> */}
        
        <form onSubmit={this.enviaDados}>
          <Legenda htmlFor="email">Usuário:</Legenda>
          <Campo ref={this.emailRef} id="email" type="text" name="email" placeholder="Usuário" required onChange={this.habilitaOuDesabilitaBotao} />
          
          <Legenda htmlFor="senha">Senha:</Legenda>
          <Campo ref={this.senhaRef} id="senha" type="password" name="senha" placeholder="Senha" required minLength={6} onChange={this.habilitaOuDesabilitaBotao} />
          
          <Botao desabilitado={this.state.desabilitado}>
            Enviar
          </Botao>
        </form>

        <Link url="/conta">Criar uma conta</Link>
        </div>
      </main>
    )
  }
}

export default withRouter(Login)