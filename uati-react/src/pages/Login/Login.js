import React, { Component } from 'react'
import Link from '../../componentes/Link/Link'
import Botao from '../../componentes/Botao/Botao'
import Legenda from '../../componentes/Legenda/Legenda'
import Campo from '../../componentes/Campo/Campo'
import * as api from '../../apis/login'
import './Login.css'
import * as storage from '../../service/helpers'

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

    let resp
    try {
      resp = await api.login(dados)
      console.log(resp)
    } catch (e) {
      alert("Usuário e senha inválido")
      return
    }

    storage.set(resp.data.token)
    this.props.history.push('/dashboard')
    
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
        {/* <p>Entre com seu email e senha.</p> */}
        
        <form onSubmit={this.enviaDados}>
          <Legenda htmlFor="email">Usuário:</Legenda>
          <Campo ref={this.emailRef} id="usuario" type="text" name="usuario" placeholder="Usuario" required onChange={this.habilitaOuDesabilitaBotao} />
          
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

export default Login