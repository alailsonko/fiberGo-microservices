import React, { Component } from 'react'
import DatePicker from 'react-datepicker'
import axios from 'axios'


export default class CreateExercises extends Component {
   
  state = {
    username: '',
    description: '',
    duration: 0,
    date: new Date(),
    users: []
  }
   
  componentDidMount() {
      axios.get('http://localhost:5000/users')
      .then(res => {
        if(res.data.length > 0){
          users: res.data.map(user => user.username),
          username: res.data[0].username

        }
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
     axios.post('http://localhost:5000/exercise/add', exercise)
     .then(res => console.log(res.data))

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