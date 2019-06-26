import React, { Component } from 'react'
import Link from '../../componentes/Link/Link'
import Botao from '../../componentes/Botao/Botao'
import Legenda from '../../componentes/Legenda/Legenda'
import Campo from '../../componentes/Campo/Campo'
import './Conta.css'
import api from '../../services/api'


class Conta extends Component {
  constructor(props) {
    super(props)

    this.nomeRef = React.createRef()
  //  this.userRef = React.createRef()
    this.usuarioRef = React.createRef()
    this.emailRef = React.createRef()
    this.senhaRef = React.createRef()

    this.state = { desabilitado: true }
  }

  habilitaOuDesabilita = () => {
    const campoNome = this.nomeRef.current
    const campoUsuario = this.usuarioRef.current
    const campoEmail = this.emailRef.current
    const campoSenha = this.senhaRef.current
    
    if (campoNome.temErro() || campoUsuario.temErro() || campoEmail.temErro() || campoSenha.temErro()) {
      this.setState({ desabilitado: true })
    } else {
      this.setState({ desabilitado: false })
    }
  
}

  cadastrar = async (e) => {
    e.preventDefault();

    const campoNome = this.nomeRef.current
    const campoUser = this.usuarioRef.current
    const campoEmail = this.emailRef.current
    const campoSenha = this.senhaRef.current

    const dados = {
      username: campoUser.getValor(),
      password: campoSenha.getValor(),
      name: campoNome.getValor(),
      email: campoEmail.getValor()
    }

    try {
      await api.post("users", dados)
    } catch (e) {
      alert("Dados inválidos")
      return
    }


    this.props.history.push('/login')

  }

  render() {
    return (
      <main>
        <form onSubmit={this.cadastrar}>
      <div className="form">
        <h1>Cadastro</h1>
        {/* <p>Envie o formulário para criar uma conta!</p> */}
        
        <Legenda htmlFor="nome">Nome:</Legenda>
        <Campo ref={this.nomeRef} id="nome" type="text" name="nome" placeholder="Nome" required minLength={10} onChange={this.habilitaOuDesabilita} />
        
        {/* <Legenda htmlFor="username">Usuário:</Legenda>
        <Campo ref={this.userRef} id="username" type="text" name="username" placeholder="Usuário" required onChange={this.habilitaOuDesabilita} /> */}
        <Legenda htmlFor="usuario">Usuário:</Legenda>
        <Campo ref={this.usuarioRef} id="usuario" type="tel" name="telefone" placeholder="Telefone" required onChange={this.habilitaOuDesabilita} />
        
        <Legenda htmlFor="email">Email:</Legenda>
        <Campo ref={this.emailRef} id="email" type="email" name="email" placeholder="Email" required onChange={this.habilitaOuDesabilita} />
        
        <Legenda htmlFor="senha">Senha:</Legenda>
        <Campo ref={this.senhaRef} id="senha" type="password" name="senha" placeholder="Senha" required minLength={6} onChange={this.habilitaOuDesabilita} />
        
        <Botao desabilitado={this.state.desabilitado}>Enviar</Botao>
  
        <Link url="/login">Fazer login</Link>
        </div>
        </form>
      </main>
    )
  }
}

export default Conta