import React from 'react'
import { useState } from 'react'
import './style.scss'
import axios from 'axios'
import { useContext } from 'react'
import { AppContext } from '../../App'
import { useNavigate } from 'react-router-dom'

const Home = () => {
  const [state, setState] = useState({
    url: "",
    custom_id: ""
  })
  const {url, custom_id} = state;

  const {onStateChange, clearState, state:{shortened_url, url_id}} = useContext(AppContext);

  const handleClick =  async () => {
    try {
      const res = await axios.post("http://localhost:8080/api/shorten_url", state);
      onStateChange(res.data);
    }catch(error) {
      console.log(error)
    }
  }

  const copyToClipboard = () => {
    window.navigator.clipboard.writeText(shortened_url);
    alert("copied")
  }

  const navigate = useNavigate()

  return (
    <div className='home'>
        <div className='form'>
            <div className='inputs'>
              <input
               type="text"
               placeholder='Enter URL' 
               id='url' 
               name="url" 
               value={url} 
               onChange={(e) =>  setState(state => ({
                ...state,
                [e.target.name]: e.target.value
              }))}/>
              <input 
               type="text" 
               placeholder='Enter custom ID for URL(optional, must be six characters long)' 
               id='custom_id' 
               name='custom_id'
               value={custom_id}
               onChange={(e) => setState(state => ({
                ...state,
                [e.target.name]: e.target.value
               }))} />
            </div>
            <button onClick={handleClick}>Shorten</button>
        </div>

        {shortened_url && <div className='display_shortened_url'>
          <span>{shortened_url}</span>
          <i className="fas fa-clipboard" onClick={copyToClipboard}></i>
          <i className="fas fa-link" onClick={() => navigate(`/${url_id}`)}></i>
          <i className="fa fa-times" aria-hidden="true" onClick={clearState}></i>
        </div>
        }
    </div>
  )
}

export default Home