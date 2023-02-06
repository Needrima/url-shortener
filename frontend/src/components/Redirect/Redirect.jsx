import axios from 'axios'
import React from 'react'
import { useEffect } from 'react'
import { useParams } from 'react-router-dom'

const Redirect = () => {
  const id = useParams()["id"]
  
  useEffect(() => {
    (async () => {
      try {
        const {data} = await axios.get(`http://localhost:8080/api/${id}`)
        window.location.replace(data.data)
      }catch(error) {
        console.log(error)
      }
    })();
  }, [id])
  return (
    <div>Redirect in a second ......</div>
  )
}

export default Redirect