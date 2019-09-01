import React, { Component } from "react";

class Message extends Component {
  constructor(props) {
    super(props);
    let temp = JSON.parse(this.props.message);

    var today = new Date();
    var date = today.getFullYear()+'-'+(today.getMonth()+1)+'-'+today.getDate();
    var time = today.getHours() + ":" + today.getMinutes() + ":" + today.getSeconds();
    var dateTime = date+' '+time;
    this.state = {
      message: temp,
      time : dateTime
    };
  }

  render() {

    return <div className="Message">
    {this.state.time} {this.state.message.body}
    </div>;
  }
}

export default Message;