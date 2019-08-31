import React, { Component } from "react";
import './App.css';
import { connect, sendMsg} from "./api";

import Header from './components/Header/Header';
import ChatHistory from './components/ChatHistory/ChatHistory';
import ChatInput from './components/ChatInput/ChatInput';

import firebase from 'firebase';
import StyledFirebaseAuth from "react-firebaseui/StyledFirebaseAuth"

firebase.initializeApp({
  apiKey:"AIzaSyA6W8G_LGv62M6O1eYGcQCr9sK4RxmdFwE",
  authDomain:"chat-challenge-49c3e.firebaseapp.com"
})

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      chatHistory: [],
      isSignedIn: false,
      }
  }

  uiConfig = {
    signinFlow: "popup",
    signInOptions: [
      firebase.auth.GoogleAuthProvider.PROVIDER_ID,
      firebase.auth.EmailAuthProvider.PROVIDER_ID
    ],
    callbacks: {
      signInSuccess: () => false
    }
  }


  send(event) {
    if(event.keyCode === 13) {
      sendMsg(event.target.value);
      event.target.value = "";
    }
  }

  componentDidMount() {
    firebase.auth().onAuthStateChanged(user => {
      this.setState({isSignedIn: !!user})
    })

    connect((msg) => {
      console.log("New Message")
      this.setState(prevState => ({
        chatHistory: [ ...this.state.chatHistory,msg],
      }));
      console.log(this.state);
    });
  }

  render() {
    return (
      <div className="App">
        {this.state.isSignedIn ? (
        <div>
          <Header />
          <ChatHistory chatHistory={this.state.chatHistory} />
          <ChatInput send={this.send} />
          <button onClick= {()=> firebase.auth().signOut()}>Sign Out</button>
        </div>
         ) :(
        <StyledFirebaseAuth
         uiConfig={this.uiConfig}
         firebaseAuth={firebase.auth()}
         />
         )}
      </div>
    );
  }
}

export default App;
