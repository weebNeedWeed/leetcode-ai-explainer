# LeetCode AI Explainer Frontend

This is the frontend for the LeetCode AI Explainer project, built with React and Tailwind CSS.

## Features

- Clean, responsive user interface
- Problem solution display with syntax highlighting
- Loading states and error handling
- Integration with the Go API backend

## Technology Stack

- React 19
- Tailwind CSS 4.1
- Vite build system
- Nginx for serving in production

## Project Structure

```
web/
├── public/              # Static assets
├── src/                 # Source code
│   ├── components/      # React components
│   ├── services/        # API service integration
│   ├── App.jsx          # Main application component
│   └── main.jsx         # Application entry point
├── index.html           # HTML template
├── package.json         # Dependencies and scripts
├── vite.config.js       # Vite configuration
├── tailwind.config.js   # Tailwind CSS configuration
├── nginx.conf           # Nginx configuration for production
└── eslint.config.js     # ESLint configuration
```

## Development

### Prerequisites

- Node.js 22+
- npm 10+

### Setup

1. Install dependencies:
   ```bash
   npm install
   ```

2. Start the development server:
   ```bash
   npm run dev
   ```

3. The app will be available at http://localhost:5173 with hot-reload enabled.

### Building for Production

```bash
npm run build
```

This will create optimized files in the `dist` directory that can be served by Nginx.

## Integration with Backend

The frontend communicates with the Go API backend through the `/api` endpoint. In development mode, API requests are proxied to the backend server running at http://localhost:9090.

## Styling

The project uses Tailwind CSS with custom theming defined in `tailwind.config.js` and extended in `main.css`.
