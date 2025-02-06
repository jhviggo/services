from PIL import Image
import sys


FILE_NAME=sys.argv[1]

with Image.open(FILE_NAME) as img:
  width, height = img.size

  # large scale optimized image
  large_size = 1300
  large_ratio = large_size/width
  if (width > large_size):
    large_image = img.resize((large_size, int(height*large_ratio)))
  else:
    large_image = img.resize((width, height))
  
  large_image.save('large-' + FILE_NAME)

  # small scale optimized image
  small_size = 600
  small_ratio = small_size/width
  small_image = img.resize((small_size, int(height*small_ratio)))
  small_image.save('small-' + FILE_NAME)
