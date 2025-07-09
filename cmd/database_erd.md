# Literary Lions Forum - Entity Relationship Diagram

## Database Schema Overview

This ERD shows the relationships between all tables in the Literary Lions Forum database.

```mermaid
erDiagram
    users {
        int id PK
        string email UK
        string username UK
        string password
        datetime created_at
    }
    
    posts {
        int id PK
        string title
        text content
        int user_id FK
        datetime created_at
        datetime updated_at
    }
    
    comments {
        int id PK
        text content
        int user_id FK
        int post_id FK
        datetime created_at
    }
    
    categories {
        int id PK
        string name UK
        string slug UK
        text description
        datetime created_at
    }
    
    post_categories {
        int id PK
        int post_id FK
        int category_id FK
        datetime created_at
    }
    
    likes {
        int id PK
        int user_id FK
        int post_id FK
        int comment_id FK
        boolean is_like
        datetime created_at
    }
    
    sessions {
        string id PK
        int user_id FK
        datetime created_at
        datetime expires_at
    }
    
    %% Relationships
    users ||--o{ posts : "creates"
    users ||--o{ comments : "writes"
    users ||--o{ likes : "gives"
    users ||--o{ sessions : "has"
    
    posts ||--o{ comments : "has"
    posts ||--o{ likes : "receives"
    posts ||--o{ post_categories : "belongs_to"
    
    comments ||--o{ likes : "receives"
    
    categories ||--o{ post_categories : "contains"