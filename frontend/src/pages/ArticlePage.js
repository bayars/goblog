import React, { useState, useEffect } from 'react';

function ArticlePage() {
  const [article, setArticle] = useState(null);

  useEffect(() => {
    const apiUrl = 'http://10.20.30.46:8080/articles/2';

    fetch(apiUrl)
      .then(response => response.json())
      .then(data => setArticle(data))
      .catch(error => console.error('Error:', error));
  }, []);



  return (
    <div className="article-page">
      {article ? (
      <div className="div">
        <div className="group">
          <div className="overlap-group">

            <div className="text-wrapper">{article[0]}</div>
            <p className="p">{article[1]}</p>
          </div>
        </div>
        <p className="feel-free-to-use">
          {article[2]}
          <br />
        </p>
      </div>
      ) : (
        <p>Loading...</p>
      )}
    </div>
  );
};
  
export default ArticlePage;
