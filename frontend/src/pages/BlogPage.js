import React, { useState, useEffect } from "react";

function BlogPage() {
  const [articles, setArticles] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchArticles = async () => {
      try {
        const response = await fetch('http://10.20.30.46:8080/articles');
        
        if (!response.ok) {
          throw new Error('Failed to fetch data');
        }

        const data = await response.json();
        setArticles(data);
      } catch (error) {
        console.error('Error:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchArticles();
  }, []);

  return (
    <div className="blog-page">
      <span className="span-2">BLOG</span>
      {loading ? (
        <p>Loading...</p>
      ) : articles.length > 0 ? (
        <>

        {articles.map((article, index) => (
          <span key={article.id || index} className="span-3">
          <div className="div-8">{article.timestamp.slice(0,10)}</div>
          <span className="span-4">
            <div className="div-9">
              {article.title}
            </div>
            <div className="div-10">
              {article.content}
            </div>
            <span className="span-5">#tag</span>
          </span>
          <br />
          </span>

      ))}
        </>

      ) : (
        <p>No articles found</p>
      )}
    </div>
  );
}

export default BlogPage;





