import React, { Component } from 'react'
import DatePicker from 'react-datepicker'



export default class CreateExercises extends Component {
   
  state = {
    username: '',
    description: '',
    duration: 0,
    date: new Date(),
    users: []
  }
   
  componentDidMount() {
    this.setState({
      users: ['test user'],
      username: 'test user'
    })
  }


  onChange(e) {
    this.setState({
          [e.target.name]: e.target.value
    })
  }
  onChangeDate(date) {
    this.setState({
      date: date
    })
  }
   onSubmit(e){
     e.preventDefault()

     const exercise = {
       username: this.state.username,
       description: this.state.description,
       duration: this.state.duration,
       date: this.state.date,
     }
     console.log(exercise);
     window.location = '/'
   }

  render(){

    return (
      <div>
         <h3>Create New Exercise Log</h3>
         <form onSubmit={this.onSubmit}>
           <label>Username:</label>
           <select ref="userInput"
            required
            className="form-control"
            value={this.state.username}
            onChange={this.onChange}
           >
             {
               this.state.users.map(function(user){
                 return <option key={user} value={user}>{user}</option>
               })
             }
           </select>
           <div className="form-group">
             <label>Description: </label>
             <input type="text"
              required
              className="form-control"
              value={this.state.description}
              onChange={this.onChange}
             />
             <label>Duration (in minutes): </label>

              <input type="text"
              required
              className="form-control"
              value={this.state.duration}
              onChange={this.onChange}
             />
             <label>Date: </label>
              <div>
              <DatePicker
              selected={this.state.date}                      
              onChange={this.onChangeDate}
             />
                </div> 
           </div>
           <div className="form-group">
             <input type="submit" value="create exercise log" className="btn btn-primary"/>
           </div>
           </form> 
      </div>
    )
  }
}