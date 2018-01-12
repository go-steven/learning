# -*- coding: utf-8 -*-

# run it with: 
#   python index.py

if __name__ == "__main__":
    import logging
    logging.basicConfig(filename='example.log', level=logging.DEBUG)
    
    logging.debug('This message should go to the log file')
    logging.info('So should this')
    logging.warning('And this, too')