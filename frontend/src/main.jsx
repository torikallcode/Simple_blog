import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import { ArticleList } from './components/ArticleList'
import { ArticleForm } from './components/ArticleForm'
import { ArticleDetail } from './components/ArticleDetail'
import axios from 'axios'

const handleCreate = (article) => {
  axios.post('http://localhost:8080/articles', article)
    .then(response => console.log(response.data))
    .catch(error => console.error(error));
};

const router = createBrowserRouter([
  {
    path: "/",
    element: <ArticleList />,
  },
  {
    path: "/articles/new",
    element: <ArticleForm onSubmit={handleCreate} />
  },
  {
    path: "/articles/:id",
    element: <ArticleDetail />
  }
])
createRoot(document.getElementById('root')).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
)
