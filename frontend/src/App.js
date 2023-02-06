import './App.css';
import React from 'react'
import {Route, Routes} from 'react-router-dom'
import Home from './components/Home/Home';
import Redirect from './components/Redirect/Redirect';
import { useState } from 'react';

export const AppContext = React.createContext();
function App() {
  const [state, setState] = useState({
    shortened_url: "",
    url_id: "",
    usageRemaining: 0,
    usage_limit_reset: 0,
  })

  const onStateChange = (data) => {
    setState(state => ({
      ...state,
      shortened_url: data.shortened_url,
      url_id: data.url_id,
      usageRemainig: data.usageRemaining,
      usage_limit_reset: data.usage_limit_reset,
    }))
  }

  const clearState = () => {
    setState({
      shortened_url: "",
      url_id: "",
      usageRemaining: 0,
      usage_limit_reset: 0,
    })
  }

  // const handle
  return (
    <AppContext.Provider value={{
      state,
      onStateChange,
      clearState
    }}>
      <Routes>
        <Route path="/" exact element={<Home />} />
        <Route path="/:id" exact element={<Redirect />} />
      </Routes>
    </AppContext.Provider>
  );
}

export default App;
