import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import FlowCanvas from './FlowCanvas'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <FlowCanvas />
    </>
  )
}

export default App
