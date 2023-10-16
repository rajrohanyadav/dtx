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
    <li>
      <a href="#generate">Generate</a>
      <ul>
        <li><a href="#uuid">UUID</a></li>
      </ul>
    </li>
  </ol>

</details>


## Encode

### Base64

```sh
dtx encode -t b64 -s user:password
```
***Output***: `dXNlcjpwYXNzd29yZA==`

## Decode

### Base64

```sh
dtx decode -t b64 -s dXNlcjpwYXNzd29yZA==
```
***Output***: `user:password`

## Generate

### UUID

```sh
dtx generate -t uuid -n 3
```
***Output***:
```
1c973675-285c-4aa2-9225-0add790b3ecd
e65bbb03-70e7-4d3d-a72e-d76cec3a383a
7c426a9f-467b-4838-accd-65cf406a6c0d
```
