import React, { useEffect, useState } from 'react'
import axios from 'axios'
import { Link } from 'react-router-dom'

export const ArticleList = () => {

  const [articles, setArticles] = useState([])

  useEffect(() => {
    axios.get('http://localhost:8080/articles')
      .then(response => setArticles(response.data))
      .catch(error => console.log(error))
  }, [])

  return (
    <div className='flex flex-col justify-center items-center pt-10'>
      <h1 className='text-3xl font-bold uppercase mb-10'>My Articles</h1>
      <ul className='grid grid-cols-2 gap-3 max-w-xl'>
        {articles.map(article => (
          <li key={article.id}>
            <Link to={`/articles/${article.id}`} className='bg-sky-500 hover:bg-sky-700 text-white cursor-pointer shadow-md font-bold py-2 px-4 h-32 w-52 rounded flex items-center justify-center'>{article.title}</Link>
          </li>
        ))}
      </ul>
    </div >
  )
}
