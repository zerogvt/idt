import unittest
from pymoticon import emotistrip


class TestDeploy(unittest.TestCase):
    def test_emotistrip(self):
        db = {
            ":)": "smiley",
            ";)": "wink"
        }
        inp = "Hello ! :)"
        want = ["Hello", ":)"]
        have = emotistrip(inp, db)
        self.assertEqual(want, have)
