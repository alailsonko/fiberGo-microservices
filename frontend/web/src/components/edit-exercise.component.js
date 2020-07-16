import React, { Component } from 'react'
import axios from 'axios'

export default class EditExercise extends Component {
  
  state = {
    username: '',
  }
   


  onChange(e) {
    this.setState({
          [e.target.name]: e.target.value
    })
  }
 
   onSubmit(e){
     e.preventDefault()

     const user = {
       username: this.state.username,
 
     }
     axios.post('http://localhost:5000/users/add', user)
     .then(res => console.log(res.data))
      .catch(err => console.log(err.error))

     console.log(user);
    this.setState({
      username: ''
    })
   }


  render(){
    return (
      <div>
       <h3>Create New user</h3>
       <form onSubmit={this.onSubmit}>
         <div className="form-control">
           <label>Username:</label>
           <input type="text"
           required
           name="username"
           value={this.state.username}
           onChange={this.onChange}
           className="form-control"/>
         </div>
         <div className="form-control">
           <input type="submit" value="create user" className="btn btn-primary"/>
         </div>
       </form>
      </div>
    )
  }
}