import unittest
from pymoticon import emotistrip, init_emo_db


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

    def test_init_emo_db(self):
        db = init_emo_db()
        self.assertTrue(':)' in db)
