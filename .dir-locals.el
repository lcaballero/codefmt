
(
 (nil . (
         (eval . (progn

                   (advice-add 'risky-local-variable-p :override #'ignore)

                   (setq default-directory
                         (locate-dominating-file buffer-file-name ".dir-locals.el"))

                   (setq compilation-read-command nil)

                   (setq compile-command
                         (format "%s%s" default-directory "run.sh build"))

                   (setq testing-command
                         (format "%s%s" default-directory "run.sh tests"))

                   )
               )
         )
   )
)

