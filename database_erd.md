# Database ERD - Literary Lions Forum

The structure of the database is visualised here:

```mermaid
erDiagram
    users {
        INTEGER user_id PK
        TEXT username
        TEXT email
        TEXT password_hash
        TIMESTAMP date_created
    }
    
    posts {
        INTEGER post_id PK
        INTEGER user_id FK
        TEXT post_title
        TEXT post_content
        TIMESTAMP post_time
        TIMESTAMP updated_at
    }
    
    comments {
        INTEGER comment_id PK
        INTEGER post_id FK
        INTEGER user_id FK
        TEXT comment_content
        TIMESTAMP comment_time
    }
    
    categories {
        INTEGER category_id PK
        TEXT name
        TEXT slug
        TEXT description
        TIMESTAMP created_at
    }
    
    post_categories {
        INTEGER id PK
        INTEGER post_id FK
        INTEGER category_id FK
        TIMESTAMP created_at
    }
    
    sessions {
        TEXT session_id PK
        INTEGER user_id FK
        TIMESTAMP start_time
        TIMESTAMP expiry_time
    }
    
    likes {
        INTEGER like_id PK
        INTEGER user_id FK
        INTEGER post_id FK
        INTEGER comment_id FK
        BOOLEAN is_like
        TIMESTAMP created_at
    }
    
    users ||--o{ posts : "creates"
    users ||--o{ comments : "writes"
    users ||--o{ sessions : "has"
    users ||--o{ likes : "gives"
    
    posts ||--o{ comments : "contains"
    posts ||--o{ likes : "receives"
    posts ||--o{ post_categories : "belongs_to"
    
    comments ||--o{ likes : "receives"
    
    categories ||--o{ post_categories : "categorizes"