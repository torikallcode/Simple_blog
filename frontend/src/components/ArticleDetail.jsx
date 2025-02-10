import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { useParams } from 'react-router-dom';

export const ArticleDetail = () => {
  const { id } = useParams();
  const [article, setArticle] = useState(null);

  useEffect(() => {
    axios.get(`http://localhost:8080/articles/${id}`)
      .then(response => setArticle(response.data))
      .catch(error => console.error(error));
  }, [id]);

  if (!article) return <div>Loading...</div>;

  return (
    <div className='flex flex-col'>
      <div className='bg-sky-500 py-28 mb-3'>
        <h1 className='text-3xl font-bold text-center text-white'>{article.title}</h1>
      </div>
      <p>{article.content}</p>
      {/* <p><em>Published on: {new Date(article.createdAt).toLocaleDateString()}</em></p> */}
    </div>
  );
};

