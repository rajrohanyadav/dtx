# Dev 10X (dtx)

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
![Build workflow][build-workflow-url]
![Coverage workflow][coverage-workflow-url]
<!-- [![LinkedIn][linkedin-shield]][linkedin-url] -->

<details>
<summary>Table of Contents</summary>

<ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#homebrew">Homebrew</a></li>
        <li><a href="#github">GitHub</a></li>
        <li><a href="#build-from-source">Build From Source</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>

</details>

## About the Project

## Built With
[![Go][go-shield]][go-url]


## Getting Started
Use any of the following ways to install `dtx`

### Homebrew
```sh
brew tap rajrohanyadav/rajrohanyadav
brew install dtx
```

### GitHub
* Go to [dtx releases](https://github.com/rajrohanyadav/dtx/releases/)
* Download the executable for you Operating system
* Add it to your `$PATH`

### Build from source

#### Prerequisites
* Go
* Cobra-cli

#### Steps to build locally
1. Clone the repo
   ```sh
   git clone https://github.com/rajrohanyadav/dtx.git
   ```
2. Build the executables
    ```sh
    make build
    ```

## Usage
_For detailed usecase and  examples, please refer to the [Documentation](https://rajrohanyadav.github.io/dtx/)_

## Roadmap
- [ ] Convert
  - [ ] JSON <> XML
  - [ ] Timestamp
  - [ ] Number base
  - [ ] cron parser
- [ ] Encoder/Decoders
  - [ ] Base64
  - [ ] JWT
- [ ] Generators
  - [ ] Hash
  - [ ] UUID
  - [ ] Lorem Ipsum
  - [ ] Checksum
- [ ] Formatters
  - [ ] JSON
  - [ ] SQL
  - [ ] XML

See the [open issues](https://github.com/rajrohanyadav/dtx/issues) for a full list of proposed features (and known issues).

## Contributing
Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are greatly appreciated.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement". Don't forget to give the project a star! Thanks again!

1. Fork the Project
1. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
1. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
1. Push to the Branch (`git push origin feature/AmazingFeature`)
1. Open a Pull Request

## License
Distributed under the Apache License. See `LICENSE.txt` for more information.

## Contact
Rohan Yadav - rajrohanyadav@gmail.com

## Acknowledgments
- I want to create a lightweight, platform independent Command Line Utility (CLI) with the tools similar to [DevToys](https://devtoys.app/).



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/rajrohanyadav/dtx.svg
[contributors-url]: https://github.com/rajrohanyadav/dtx/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/rajrohanyadav/dtx.svg
[forks-url]: https://github.com/rajrohanyadav/dtx/network/members
[stars-shield]: https://img.shields.io/github/stars/rajrohanyadav/dtx.svg
[stars-url]: https://github.com/rajrohanyadav/dtx/stargazers
[issues-shield]: https://img.shields.io/github/issues/rajrohanyadav/dtx.svg
[issues-url]: https://github.com/rajrohanyadav/dtx/issues
[license-shield]: https://img.shields.io/github/license/rajrohanyadav/dtx.svg
[license-url]: https://github.com/rajrohanyadav/dtx/blob/master/LICENSE.txt
[build-workflow]: Build
[build-workflow-url]: https://github.com/rajrohanyadav/dtx/actions/workflows/build.yml/badge.svg
[coverage-workflow]: Coverage
[coverage-workflow-url]: https://github.com/rajrohanyadav/dtx/actions/workflows/coverage.yml/badge.svg
<!-- [linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/rajrohanyadav -->
<!-- [product-screenshot]: images/screenshot.png -->
[go-shield]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[go-url]: https://go.dev/
