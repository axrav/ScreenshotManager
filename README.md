# ScreenshotManager
A simple screenshot manager for my personal us as i am lazy to copy paste screenshots everytime

## Copy example.env to .env and add your values.

## You must need Go 1.17 atleast to run this

### Use a Post request to post the screenshot, it returns a url of the uploaded screenshot

### Via Curl and Maim:
```maim -sqm 10 | curl -sL --header "password: YOURPASS" -F file=@- hostname/upload | tr -d '"' | xclip -selection clipboard  ```

### Example
``` maim -sqm 10 | curl -sL --header "password: vegeta" -F file=@- https://vegetaxd.me/upload | tr -d '"' | xclip -selection clipboard ```

Thanks to [TheHamkerCat](https://github.com/thehamkercat) for the idea xD
