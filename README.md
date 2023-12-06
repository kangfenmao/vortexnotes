# VortexNotes

VortexNotes is a unique web-based note-taking software that offers a different approach to managing your notes.

## Overview

<a href="" target="_blank">
  <div align="center">
    <img width="668" alt="image" src="https://github.com/kangfenmao/vortexnotes/assets/8253512/6c7e6893-3e41-4580-90a3-747cafba0638">
  </div>
</a>

## Introduction

VortexNotes is designed to be a simple yet powerful note-taking tool. It provides users with a single entry point to submit new content without the need for note organization or management features. It serves as a "black hole" for your notes, where you can simply paste your content and let it disappear into the void.

## Features

- **Submit and Forget**: With VortexNotes, there is no need to worry about organizing or managing your notes. Simply paste your content into the application, and it will be absorbed into the void, never to be seen again.

- **Effortless Searching**: VortexNotes offers a search interface that allows you to quickly find specific notes. Despite the lack of note management, the powerful search feature enables you to retrieve your desired content effortlessly.

- **Clean and Minimalistic**: The user interface of VortexNotes is designed to be clean and minimalistic, focusing solely on the note submission and search functionalities. This ensures a distraction-free environment for your note-taking.

## Getting Started

To start using VortexNotes, follow these steps:

1. Clone this repository.
2. Install the necessary dependencies.
3. Run the application.
4. Access the application via the provided URL in your web browser.
5. Paste your content into the submission field and let it disappear into the void.
6. Use the search interface to retrieve your notes when needed.

## Docker compose

```yml
version: '3'

services:
  vortexnotes:
    container_name: vortexnotes
    image: vortexnotes:latest
    environment:
      VORTEXNOTES_PASSCODE: fcewNEjioVf8QlTHzV4U
      VORTEXNOTES_AUTH_SCOPE: show,create,edit,delete
    ports:
      - "7701:7701"
    volumes:
      - ./app/data/notes:/data/notes
      - ./app/data/vortexnotes:/data/vortexnotes
```

If you only require a password to safeguard the writing actions.

```yml
environment:
  VORTEXNOTES_PASSCODE: fcewNEjioVf8QlTHzV4U
  VORTEXNOTES_AUTH_SCOPE: create,edit,delete
```

If you want anyone to be able to view and edit. You can remove `VORTEXNOTES_PASSCODE` and `VORTEXNOTES_AUTH_SCOPE` env.

## Service Port

| Service | Port | Public |
|---------|------|--------|
| API     | 7701 | true   |
| Web     | 7701 | true   |

## Contributing

Contributions are welcome! If you have any ideas, suggestions, or bug reports, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
