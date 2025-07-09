flowchart LR
    subgraph UsersBox [" "]
        users["<div style='background:#4A90E2;color:white;padding:5px;border-radius:3px;'><b>users</b></div>
        <div style='padding:10px;'>
        INTEGER user_id PK<br/>
        TEXT username<br/>
        TEXT email<br/>
        TEXT password_hash<br/>
        TIMESTAMP date_created
        </div>"]
    end
    
    subgraph PostsBox [" "]
        posts["<div style='background:#4A90E2;color:white;padding:5px;border-radius:3px;'><b>posts</b></div>
        <div style='padding:10px;'>
        INTEGER post_id PK<br/>
        INTEGER user_id FK<br/>
        TEXT post_title<br/>
        TEXT post_content<br/>
        TIMESTAMP post_time<br/>
        TIMESTAMP updated_at
        </div>"]
    end
    
    subgraph CommentsBox [" "]
        comments["<div style='background:#4A90E2;color:white;padding:5px;border-radius:3px;'><b>comments</b></div>
        <div style='padding:10px;'>
        INTEGER comment_id PK<br/>
        INTEGER post_id FK<br/>
        INTEGER user_id FK<br/>
        TEXT comment_content<br/>
        TIMESTAMP comment_time
        </div>"]
    end
    
    subgraph CategoriesBox [" "]
        categories["<div style='background:#4A90E2;color:white;padding:5px;border-radius:3px;'><b>categories</b></div>
        <div style='padding:10px;'>
        INTEGER category_id PK<br/>
        TEXT name<br/>
        TEXT slug<br/>
        TEXT description<br/>
        TIMESTAMP created_at
        </div>"]
    end
    
    subgraph PostCategoriesBox [" "]
        post_categories["<div style='background:#4A90E2;color:white;padding:5px;border-radius:3px;'><b>post_categories</b></div>
        <div style='padding:10px;'>
        INTEGER id PK<br/>
        INTEGER post_id FK<br/>
        INTEGER category_id FK<br/>
        TIMESTAMP created_at
        </div>"]
    end
    
    subgraph SessionsBox [" "]
        sessions["<div style='background:#4A90E2;color:white;padding:5px;border-radius:3px;'><b>sessions</b></div>
        <div style='padding:10px;'>
        TEXT session_id PK<br/>
        INTEGER user_id FK<br/>
        TIMESTAMP start_time<br/>
        TIMESTAMP expiry_time
        </div>"]
    end
    
    subgraph LikesBox [" "]
        likes["<div style='background:#4A90E2;color:white;padding:5px;border-radius:3px;'><b>likes</b></div>
        <div style='padding:10px;'>
        INTEGER like_id PK<br/>
        INTEGER user_id FK<br/>
        INTEGER post_id FK<br/>
        INTEGER comment_id FK<br/>
        BOOLEAN is_like<br/>
        TIMESTAMP created_at
        </div>"]
    end
    
    %% Relationships with better arrows
    users -.->|creates| posts
    users -.->|writes| comments
    users -.->|has| sessions
    users -.->|gives| likes
    
    posts -.->|contains| comments
    posts -.->|receives| likes
    posts -.->|belongs to| post_categories
    
    comments -.->|receives| likes
    
    categories -.->|categorizes| post_categories
    
    %% Styling
    linkStyle 0 stroke:#2E5C8A,stroke-width:2px
    linkStyle 1 stroke:#2E5C8A,stroke-width:2px
    linkStyle 2 stroke:#2E5C8A,stroke-width:2px
    linkStyle 3 stroke:#2E5C8A,stroke-width:2px
    linkStyle 4 stroke:#2E5C8A,stroke-width:2px
    linkStyle 5 stroke:#2E5C8A,stroke-width:2px
    linkStyle 6 stroke:#2E5C8A,stroke-width:2px
    linkStyle 7 stroke:#2E5C8A,stroke-width:2px
    linkStyle 8 stroke:#2E5C8A,stroke-width:2px