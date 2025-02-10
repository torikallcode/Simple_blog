import React, { useState } from 'react';

export const ArticleForm = ({ article, onSubmit }) => {
  const [title, setTitle] = useState(article ? article.title : '');
  const [content, setContent] = useState(article ? article.content : '');

  const handleSubmit = (e) => {
    e.preventDefault();
    const newArticle = { title, content };
    onSubmit(newArticle);
    setTitle('');
    setContent('');
  };

  return (
    <div className='flex flex-col pt-10 justify-center items-center'>
      <h1 className='text-2xl font-bold uppercase text-center mb-10'>Create Your Article</h1>
      <form onSubmit={handleSubmit} className='flex flex-col space-y-3'>
        <input
          type="text"
          placeholder="Title"
          value={title}
          className="w-[300px] border border-slate-200 rounded-lg py-3 px-5 outline-none  bg-transparent"

          onChange={(e) => setTitle(e.target.value)}
        />
        <textarea
          placeholder="Content"
          value={content}
          className="w-[300px] border border-slate-200 rounded-lg py-3 px-5 outline-none  bg-transparent"

          onChange={(e) => setContent(e.target.value)}
        />
        <button
          type="submit"
          className='bg-sky-500 hover:bg-sky-700 text-white cursor-pointer shadow-md font-bold py-2 px-4 rounded'
        >Save</button>
      </form>
    </div>
  );
};

