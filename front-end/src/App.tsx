import { Toaster } from '@/components/ui/sonner'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import LoginPage from './pages/Login/LoginPage';
import HomePage from './pages/Home/HomePage';
import RegisterPage from './pages/Login/RegisterPage';

function App() {
  return (
    <Router>
      <Routes>
        <Route path='/login' element={<LoginPage/>} />
        <Route path='/register' element={<RegisterPage/>}/>
        <Route path='/' element={<HomePage/>}/>
      </Routes>
      <Toaster theme='dark' richColors/>
    </Router>
  )
}

export default App
