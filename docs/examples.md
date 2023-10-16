# Examples

<details>
<summary>Table of Contents</summary>

<ol>
    <li>
      <a href="#encode">Encode</a>
      <ul>
        <li><a href="#base64">Base 64</a></li>
      </ul>
    </li>
    <li>
      <a href="#decode">Decode</a>
      <ul>
        <li><a href="#base64">Base 64</a></li>
      </ul>
    </li>
  </ol>

</details>


## Encode

### Base64

```sh
> dtx encode -t b64 -s user:password
dXNlcjpwYXNzd29yZA==
```

## Decode

### Base64

```sh
> dtx decode -t b64 -s dXNlcjpwYXNzd29yZA==
user:password
```
